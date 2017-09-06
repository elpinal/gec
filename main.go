package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/elpinal/gec/ast"
	"github.com/elpinal/gec/parser"
	"github.com/elpinal/gec/token"
	"github.com/elpinal/gec/types"

	"llvm.org/llvm/bindings/go/llvm"

	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
)

func main() {
	logFile := flag.String("log", "", "specify `filename` to output LLVM IR")
	printIR := flag.Bool("printir", false, "print IR")
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stdout, "gec: no Elacht source file given")
		os.Exit(1)
	}

	err := runMain(flag.Arg(0), *printIR, logFile)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

func runMain(filename string, printIR bool, logFile *string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Wrap(err, "gec")
	}
	return run(b, filename, printIR, logFile)
}

type Builder struct {
	llvm.Builder
	module llvm.Module
	env    map[string]types.Expr
	params map[string]llvm.Value
	decls  map[string]*ast.Decl
	refers map[string][]string
	entry  llvm.BasicBlock
}

func newBuilder(lb llvm.Builder) *Builder {
	return &Builder{
		Builder: lb,
		env:     make(map[string]types.Expr),
		params:  make(map[string]llvm.Value),
		decls:   make(map[string]*ast.Decl),
		refers:  make(map[string][]string),
	}
}

func run(input []byte, filename string, printIR bool, logFile *string) error {
	builder := newBuilder(llvm.NewBuilder())
	builder.module = llvm.NewModule(filename)

	main := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	llvm.AddFunction(builder.module, "main", main)
	block := llvm.AddBasicBlock(builder.module.NamedFunction("main"), "entry")
	builder.entry = block
	builder.SetInsertPoint(block, block.FirstInstruction())

	decls, err := parser.Parse(input)
	if err != nil {
		return err
	}
	expr, err := builder.reserve(decls)
	if err != nil {
		return err
	}
	a, err := builder.genIR(expr, "")
	if err != nil {
		return err
	}

	if printIR {
		pp.Println(a)
	}
	v, err := builder.gen(a, &types.TInt{})
	if err != nil {
		return err
	}
	ti := types.TI{}
	t, err := ti.TypeInference(types.TypeEnv{}, a)
	if err != nil {
		return err
	}
	switch t.(type) {
	case *types.TInt:
	default:
		return fmt.Errorf("expected int, got %v", t)
	}
	builder.CreateRet(v)

	if logFile != nil {
		ioutil.WriteFile(*logFile, []byte(builder.module.String()), 0666)
	}

	if err := llvm.VerifyModule(builder.module, llvm.ReturnStatusAction); err != nil {
		return err
	}

	engine, err := llvm.NewExecutionEngine(builder.module)
	if err != nil {
		return err
	}

	funcResult := engine.RunFunction(builder.module.NamedFunction("main"), []llvm.GenericValue{})
	fmt.Println(funcResult.Int(false))
	return nil
}

func (b *Builder) reserve(wd *ast.WithDecls) (ast.Expr, error) {
	for _, decl := range wd.Decls {
		if prev, found := b.decls[decl.LName()]; found {
			return nil, fmt.Errorf("redeclared at %v: %q (previously declared at %v)", decl.Pos(), decl.LName(), prev.Pos())
		}
		b.decls[decl.LName()] = decl
	}
	return wd.Expr, nil
}

func (b *Builder) resolve(tok token.Token) (types.Expr, error) {
	decl, found := b.decls[tok.Lit]
	if !found {
		return nil, fmt.Errorf("%v: unknown name: %q", tok.Position, tok.Lit)
	}
	t, err := b.genDecl(decl)
	if err != nil {
		return nil, err
	}
	b.env[tok.Lit] = t
	return t, nil
}

func (b *Builder) genDecl(decl *ast.Decl) (types.Expr, error) {
	v, err := b.genIR(decl.RHS, decl.LHS.Lit)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (b *Builder) checkCR(name, referredFrom string) error {
	for _, r := range b.refers[name] {
		if r == referredFrom {
			return fmt.Errorf("circular reference: %s", r)
		}
		err := b.checkCR(r, referredFrom)
		if err != nil {
			return err
		}
	}
	return nil
}

type Value struct {
	v llvm.Value
	t types.Type
}

func (b *Builder) genIR(expr ast.Expr, referredFrom string) (types.Expr, error) {
	switch x := expr.(type) {
	case *ast.Ident:
		if x.Name.Lit == referredFrom {
			return nil, fmt.Errorf("%v: self-reference: %q", x.Name.Position, x.Name.Lit)
		}
		// Note that there is possibility of duplication.
		b.refers[referredFrom] = append(b.refers[referredFrom], x.Name.Lit)
		err := b.checkCR(x.Name.Lit, referredFrom)
		if err != nil {
			return nil, err
		}
		t, found := b.env[x.Name.Lit]
		if !found {
			t, err = b.resolve(x.Name)
			if err != nil {
				return nil, err
			}
		}
		return t, nil
	case *ast.App:
		e1, err := b.genIR(x.Fn, referredFrom)
		if err != nil {
			return nil, err
		}
		e2, err := b.genIR(x.Arg, referredFrom)
		if err != nil {
			return nil, err
		}
		return &types.EApp{e1, e2}, nil
	case *ast.Abs:
		o, ok := b.env[x.Param.Lit]
		b.env[x.Param.Lit] = &types.EVar{x.Param.Lit}
		e, err := b.genIR(x.Body, referredFrom)
		if err != nil {
			return nil, err
		}
		delete(b.env, x.Param.Lit)
		if ok {
			b.env[x.Param.Lit] = o
		}
		return &types.EAbs{x.Param.Lit, e}, nil
	case *ast.Int:
		n, err := strconv.Atoi(x.X.Lit)
		if err != nil {
			return nil, err
		}
		return &types.EInt{n}, nil
	case *ast.Bool:
		switch x.X.Lit {
		case "true":
			return &types.EBool{true}, nil
		case "false":
			return &types.EBool{false}, nil
		}
		return nil, fmt.Errorf("invalid boolean value: %v", x)
	case *ast.Add:
		e1, err := b.genIR(x.X, referredFrom)
		if err != nil {
			return nil, err
		}
		e2, err := b.genIR(x.Y, referredFrom)
		if err != nil {
			return nil, err
		}
		return &types.EArithBinOp{types.Add, e1, e2}, nil
	case *ast.Sub:
		e1, err := b.genIR(x.X, referredFrom)
		if err != nil {
			return nil, err
		}
		e2, err := b.genIR(x.Y, referredFrom)
		if err != nil {
			return nil, err
		}
		return &types.EArithBinOp{types.Sub, e1, e2}, nil
	case *ast.Mul:
		e1, err := b.genIR(x.X, referredFrom)
		if err != nil {
			return nil, err
		}
		e2, err := b.genIR(x.Y, referredFrom)
		if err != nil {
			return nil, err
		}
		return &types.EArithBinOp{types.Mul, e1, e2}, nil
	case *ast.Div:
		e1, err := b.genIR(x.X, referredFrom)
		if err != nil {
			return nil, err
		}
		e2, err := b.genIR(x.Y, referredFrom)
		if err != nil {
			return nil, err
		}
		return &types.EArithBinOp{types.Div, e1, e2}, nil
	case *ast.If:
		cond, err := b.genIR(x.Cond, referredFrom)
		if err != nil {
			return nil, err
		}
		e1, err := b.genIR(x.E1, referredFrom)
		if err != nil {
			return nil, err
		}
		e2, err := b.genIR(x.E2, referredFrom)
		if err != nil {
			return nil, err
		}
		return &types.EIf{cond, e1, e2}, nil
	case *ast.Cmp:
		e1, err := b.genIR(x.LHS, referredFrom)
		if err != nil {
			return nil, err
		}
		e2, err := b.genIR(x.RHS, referredFrom)
		if err != nil {
			return nil, err
		}
		var op types.CmpOp
		switch x.Op {
		case ast.Eq:
			op = types.Eq
		case ast.NE:
			op = types.NE
		case ast.LT:
			op = types.LT
		case ast.GT:
			op = types.GT
		case ast.LE:
			op = types.LE
		case ast.GE:
			op = types.GE
		default:
			return nil, fmt.Errorf("unsupported comparison: %v", expr)
		}
		return &types.ECmp{op, e1, e2}, nil
	case *ast.NilList:
		return &types.ENil{}, nil
	case *ast.ParenExpr:
		return b.genIR(x.X, referredFrom)
	}
	return nil, fmt.Errorf("unknown expression: %[1]v (type: %[1]T)", expr)
}

func (b *Builder) gen(expr types.Expr, expected types.Type) (llvm.Value, error) {
	switch x := expr.(type) {
	case *types.EApp:
		ti := types.TI{}
		t, err := ti.TypeInference(types.TypeEnv{}, x.Arg)
		if err != nil {
			return llvm.Value{}, err
		}

		a, err := b.gen(x.Fn, &types.TFun{t, expected})
		if err != nil {
			return llvm.Value{}, err
		}

		arg, err := b.gen(x.Arg, t)
		if err != nil {
			return llvm.Value{}, err
		}
		return b.CreateCall(a, []llvm.Value{arg}, "call"), nil
	case *types.EAbs:
		f := llvm.FunctionType(
			llvmType(expected.(*types.TFun).Body),
			[]llvm.Type{llvmType(expected.(*types.TFun).Arg)},
			false,
		)
		v := llvm.AddFunction(b.module, "fun", f)
		v.Param(0).SetName(x.Param)
		block := llvm.AddBasicBlock(v, "entry")
		prevEntry := b.entry
		b.entry = block
		b.SetInsertPointAtEnd(block)

		b.params[x.Param] = v.Param(0)
		a, err := b.gen(x.Body, expected.(*types.TFun).Body)
		if err != nil {
			return llvm.Value{}, err
		}

		b.CreateRet(a)
		b.SetInsertPointAtEnd(prevEntry)
		b.entry = prevEntry
		return v, err
	case *types.EInt:
		return llvm.ConstInt(llvm.Int32Type(), uint64(x.Value), false), nil
	case *types.EBool:
		// 0 for false, 1 for true.
		var n int
		if x.Value {
			n = 1
		}
		return llvm.ConstInt(llvm.Int1Type(), uint64(n), false), nil
	case *types.EVar:
		v, ok := b.params[x.Name]
		if !ok {
			return llvm.Value{}, fmt.Errorf("gen: unbound variable: %v", expr)
		}
		return v, nil
	case *types.EArithBinOp:
		v1, err := b.gen(x.E1, &types.TInt{})
		if err != nil {
			return llvm.Value{}, err
		}
		v2, err := b.gen(x.E2, &types.TInt{})
		if err != nil {
			return llvm.Value{}, err
		}
		switch x.Op {
		case types.Add:
			return b.CreateAdd(v1, v2, "add"), nil
		case types.Sub:
			return b.CreateSub(v1, v2, "sub"), nil
		case types.Mul:
			return b.CreateMul(v1, v2, "mul"), nil
		case types.Div:
			return b.CreateUDiv(v1, v2, "div"), nil
		}
	case *types.EIf:
		cond, err := b.gen(x.Cond, &types.TBool{})
		if err != nil {
			return llvm.Value{}, err
		}
		f := b.entry.Parent()
		b1 := llvm.AddBasicBlock(f, "then")
		b2 := llvm.AddBasicBlock(f, "else")
		bEnd := llvm.AddBasicBlock(f, "end")
		b.CreateCondBr(cond, b1, b2)

		b.SetInsertPointAtEnd(b1)
		v1, err := b.gen(x.E1, expected)
		if err != nil {
			return llvm.Value{}, err
		}
		b.CreateBr(bEnd)

		b.SetInsertPointAtEnd(b2)
		v2, err := b.gen(x.E2, expected)
		if err != nil {
			return llvm.Value{}, err
		}
		b.CreateBr(bEnd)

		b.SetInsertPointAtEnd(bEnd)
		phi := b.CreatePHI(llvmType(expected), "phi")
		phi.AddIncoming([]llvm.Value{v1, v2}, []llvm.BasicBlock{b1, b2})
		return phi, nil
	case *types.ECmp:
		ti := types.TI{}
		t, err := ti.TypeInference(types.TypeEnv{}, x.E1)
		if err != nil {
			return llvm.Value{}, err
		}
		lhs, err := b.gen(x.E1, t)
		if err != nil {
			return llvm.Value{}, err
		}
		rhs, err := b.gen(x.E2, t)
		if err != nil {
			return llvm.Value{}, err
		}
		switch t.(type) {
		case *types.TInt:
			switch x.Op {
			case types.Eq:
				return b.CreateICmp(llvm.IntEQ, lhs, rhs, "eq"), nil
			case types.NE:
				return b.CreateICmp(llvm.IntNE, lhs, rhs, "ne"), nil
			case types.LT:
				return b.CreateICmp(llvm.IntULT, lhs, rhs, "lt"), nil
			case types.GT:
				return b.CreateICmp(llvm.IntUGT, lhs, rhs, "gt"), nil
			case types.LE:
				return b.CreateICmp(llvm.IntULE, lhs, rhs, "le"), nil
			case types.GE:
				return b.CreateICmp(llvm.IntUGE, lhs, rhs, "ge"), nil
			}
		case *types.TBool:
			switch x.Op {
			case types.Eq:
				return b.CreateICmp(llvm.IntEQ, lhs, rhs, "eq"), nil
			case types.NE:
				return b.CreateICmp(llvm.IntNE, lhs, rhs, "ne"), nil
			}
		}
		return llvm.Value{}, fmt.Errorf("unsupported comparison: %#v", expr)
	case *types.ENil:
		return llvm.ConstArray(llvmType(expected), []llvm.Value{}), nil
	}
	return llvm.Value{}, fmt.Errorf("LLVM IR generation: unexpected expression: %#v", expr)
}

func llvmType(t types.Type) llvm.Type {
	switch x := t.(type) {
	case *types.TInt:
		return llvm.Int32Type()
	case *types.TBool:
		return llvm.Int1Type()
	case *types.TFun:
		return llvm.FunctionType(llvmType(x.Body), []llvm.Type{llvmType(x.Arg)}, false)
	}
	panic(fmt.Sprintf("converting type to LLVM's one: unexpected error: %#v", t))
}
