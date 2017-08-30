package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/elpinal/gec/ast"
	"github.com/elpinal/gec/token"

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
	env    map[string]Value
	decls  map[string]*ast.Decl
	refers map[string][]string
	entry  llvm.BasicBlock
}

func newBuilder(lb llvm.Builder) *Builder {
	return &Builder{
		Builder: lb,
		env:     make(map[string]Value),
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
	a, err := builder.gen(expr, "")
	if err != nil {
		return err
	}

	builder.CreateRet(a.v)

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

func (b *Builder) resolve(tok token.Token) (Value, error) {
	decl, found := b.decls[tok.Lit]
	if !found {
		return Value{}, fmt.Errorf("%v: unknown name: %q", tok.Position, tok.Lit)
	}
	t, err := b.genDecl(decl)
	if err != nil {
		return Value{}, err
	}
	b.env[tok.Lit] = t
	return t, nil
}

func (b *Builder) genDecl(decl *ast.Decl) (Value, error) {
	v, err := b.gen(decl.RHS, decl.LHS.Lit)
	if err != nil {
		return Value{}, err
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

func (b *Builder) gen(expr ast.Expr, referredFrom string) (Value, error) {
	switch x := expr.(type) {
	case *ast.Ident:
		if x.Name.Lit == referredFrom {
			return Value{}, fmt.Errorf("%v: self-reference: %q", x.Name.Position, x.Name.Lit)
		}
		// Note that there is possibility of duplication.
		b.refers[referredFrom] = append(b.refers[referredFrom], x.Name.Lit)
		err := b.checkCR(x.Name.Lit, referredFrom)
		if err != nil {
			return Value{}, err
		}
		t, found := b.env[x.Name.Lit]
		if !found {
			t, err = b.resolve(x.Name)
			if err != nil {
				return Value{}, err
			}
		}
		return t, nil
	case *ast.App:
		var err error
		t, found := b.env[x.FnName.Lit]
		if !found {
			t, err = b.resolve(x.FnName)
			if err != nil {
				return Value{}, err
			}
		}
		args := make([]llvm.Value, len(x.Args))
		for i, arg := range x.Args {
			v, err := b.gen(arg, referredFrom)
			if err != nil {
				return Value{}, err
			}
			args[i] = v.v
		}
		return Value{v: b.CreateCall(t.v, args, "call")}, nil
	case *ast.Int:
		n, err := strconv.Atoi(x.X.Lit)
		if err != nil {
			return Value{}, err
		}
		return Value{
			v: llvm.ConstInt(llvm.Int32Type(), uint64(n), false),
			t: &types.TInt{},
		}, nil
	case *ast.Add:
		v1, err := b.gen(x.X, referredFrom)
		if err != nil {
			return Value{}, err
		}
		v2, err := b.gen(x.Y, referredFrom)
		if err != nil {
			return Value{}, err
		}
		return Value{
			v: b.CreateAdd(v1.v, v2.v, "add"),
			t: &types.TInt{},
		}, nil
	case *ast.Sub:
		v1, err := b.gen(x.X, referredFrom)
		if err != nil {
			return Value{}, err
		}
		v2, err := b.gen(x.Y, referredFrom)
		if err != nil {
			return Value{}, err
		}
		return Value{
			v: b.CreateSub(v1.v, v2.v, "sub"),
			t: &types.TInt{},
		}, nil
	case *ast.Mul:
		v1, err := b.gen(x.X, referredFrom)
		if err != nil {
			return Value{}, err
		}
		v2, err := b.gen(x.Y, referredFrom)
		if err != nil {
			return Value{}, err
		}
		return Value{
			v: b.CreateMul(v1.v, v2.v, "mul"),
			t: &types.TInt{},
		}, nil
	case *ast.Div:
		v1, err := b.gen(x.X, referredFrom)
		if err != nil {
			return Value{}, err
		}
		v2, err := b.gen(x.Y, referredFrom)
		if err != nil {
			return Value{}, err
		}
		return Value{
			v: b.CreateUDiv(v1.v, v2.v, "div"),
			t: &types.TInt{},
		}, nil
	}
	return Value{}, fmt.Errorf("unknown expression: %v", expr)
}
