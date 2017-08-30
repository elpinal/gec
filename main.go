package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/elpinal/gec/ast"
	"github.com/elpinal/gec/token"

	"github.com/k0kubun/pp"

	"llvm.org/llvm/bindings/go/llvm"

	"github.com/elpinal/types-go"
)

func main() {
	logFile := flag.String("log", "", "specify `filename` to output LLVM IR")
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stdout, "gec: no Elacht source file given")
		os.Exit(1)
	}
	b, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stdout, "gec: %v\n", err)
		os.Exit(1)
	}
	err = run(b, logFile)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

type Builder struct {
	llvm.Builder
	module llvm.Module
	env    map[string]types.Expr
	decls  map[string]*ast.Decl
	refers map[string][]string
	entry  llvm.BasicBlock
}

func newBuilder(lb llvm.Builder) *Builder {
	return &Builder{
		Builder: lb,
		env:     make(map[string]types.Expr),
		decls:   make(map[string]*ast.Decl),
		refers:  make(map[string][]string),
	}
}

func run(input []byte, logFile *string) error {
	builder := newBuilder(llvm.NewBuilder())
	builder.module = llvm.NewModule("gec")

	main := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	llvm.AddFunction(builder.module, "main", main)
	block := llvm.AddBasicBlock(builder.module.NamedFunction("main"), "entry")
	builder.entry = block
	builder.SetInsertPoint(block, block.FirstInstruction())

	decls, err := parse(input)
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

	pp.Println(a)
	v, err := builder.gen(a)
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

	if err := llvm.VerifyModule(builder.module, llvm.ReturnStatusAction); err != nil {
		return err
	}
	if logFile != nil {
		ioutil.WriteFile(*logFile, []byte(builder.module.String()), 0666)
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
		e, err := b.genIR(x.Body, referredFrom)
		if err != nil {
			return nil, err
		}
		return &types.EAbs{x.Param.Lit, e}, nil
	case *ast.Int:
		n, err := strconv.Atoi(x.X.Lit)
		if err != nil {
			return nil, err
		}
		return &types.EInt{n}, nil
	case *ast.Add:
		e1, err := b.genIR(x.X, referredFrom)
		if err != nil {
			return nil, err
		}
		e2, err := b.genIR(x.Y, referredFrom)
		if err != nil {
			return nil, err
		}
		return &ArithBinOp{EAdd, e1, e2}, nil
	case *ast.Sub:
		e1, err := b.genIR(x.X, referredFrom)
		if err != nil {
			return nil, err
		}
		e2, err := b.genIR(x.Y, referredFrom)
		if err != nil {
			return nil, err
		}
		return &ArithBinOp{ESub, e1, e2}, nil
	case *ast.Mul:
		e1, err := b.genIR(x.X, referredFrom)
		if err != nil {
			return nil, err
		}
		e2, err := b.genIR(x.Y, referredFrom)
		if err != nil {
			return nil, err
		}
		return &ArithBinOp{EMul, e1, e2}, nil
	case *ast.Div:
		e1, err := b.genIR(x.X, referredFrom)
		if err != nil {
			return nil, err
		}
		e2, err := b.genIR(x.Y, referredFrom)
		if err != nil {
			return nil, err
		}
		return &ArithBinOp{EDiv, e1, e2}, nil
	}
	return nil, fmt.Errorf("unknown expression: %v", expr)
}

func (b *Builder) gen(expr types.Expr) (llvm.Value, error) {
	switch x := expr.(type) {
	case *types.EApp:
		ti := types.TI{}
		t, err := ti.TypeInference(types.TypeEnv{}, x.Arg)
		if err != nil {
			return llvm.Value{}, err
		}
		f := llvm.FunctionType(
			llvm.Int32Type(),
			[]llvm.Type{llvmType(t)},
			false,
		)
		v := llvm.AddFunction(b.module, "fun", f)
		block := llvm.AddBasicBlock(v, "entry")
		b.SetInsertPointAtEnd(block)

		_, err = b.gen(x.Fn)
		if err != nil {
			return llvm.Value{}, err
		}

		arg, err := b.gen(x.Arg)
		if err != nil {
			return llvm.Value{}, err
		}
		return b.CreateCall(v, []llvm.Value{arg}, "call"), nil
	case *types.EAbs:
		a, err := b.gen(x.Body)
		if err != nil {
			return llvm.Value{}, err
		}

		b.CreateRet(a)
		b.SetInsertPointAtEnd(b.entry)
		return llvm.Value{}, err
	case *types.EInt:
		return llvm.ConstInt(llvm.Int32Type(), uint64(x.Value), false), nil
	}
	return llvm.Value{}, fmt.Errorf("gen: unexpected type: %#v", expr)
}

func llvmType(t types.Type) llvm.Type {
	switch t.(type) {
	case *types.TInt:
		return llvm.Int32Type()
	}
	panic(fmt.Sprintf("converting type to LLVM's one: unexpected error: %v", t))
}
