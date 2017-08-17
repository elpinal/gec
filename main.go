package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/elpinal/gec/ast"

	"llvm.org/llvm/bindings/go/llvm"
)

func main() {
	logFile := flag.String("log", "", "specify `filename` to output LLVM IR")
	flag.Parse()
	if flag.NArg() < 1 {
		return
	}
	err := run([]byte(flag.Arg(0)), logFile)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

type Builder struct {
	llvm.Builder
	module llvm.Module
	env    map[string]llvm.Value
	decls  map[string]ast.Decl
	refers map[string][]string
	entry  llvm.BasicBlock
}

func newBuilder(lb llvm.Builder) *Builder {
	return &Builder{
		Builder: lb,
		env:     make(map[string]llvm.Value),
		decls:   make(map[string]ast.Decl),
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

	builder.CreateRet(a)

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
		if _, found := b.decls[decl.LName()]; found {
			//TODO: Show previously declared position.
			return nil, fmt.Errorf("redeclared: %s", decl.LName())
		}
		b.decls[decl.LName()] = decl
	}
	return wd.Expr, nil
}

func (b *Builder) resolve(name string) (llvm.Value, error) {
	decl, found := b.decls[name]
	if !found {
		return llvm.Value{}, fmt.Errorf("unknown name: %s", name)
	}
	t, err := b.genDecl(decl)
	if err != nil {
		return llvm.Value{}, err
	}
	b.env[name] = t
	return t, nil
}

func (b *Builder) genDecl(decl ast.Decl) (llvm.Value, error) {
	switch x := decl.(type) {
	case *ast.Assign:
		v, err := b.gen(x.RHS, x.LHS.Lit)
		if err != nil {
			return llvm.Value{}, err
		}
		return v, nil
	case *ast.DeclFunc:
		params := make([]llvm.Type, len(x.Args))
		for i := range x.Args {
			params[i] = llvm.Int32Type()
		}
		f := llvm.FunctionType(llvm.Int32Type(), params, false)
		v := llvm.AddFunction(b.module, x.Name.Lit, f)
		for i, name := range x.Args {
			v.Param(i).SetName(name.Lit)
		}
		block := llvm.AddBasicBlock(v, "entry")
		b.SetInsertPointAtEnd(block)

		topEnv := make(map[string]llvm.Value, len(b.env))
		for k, v := range b.env {
			topEnv[k] = v
		}
		for i, name := range x.Args {
			b.env[name.Lit] = v.Param(i)
		}
		ret, err := b.gen(x.RHS, x.Name.Lit)
		if err != nil {
			return llvm.Value{}, err
		}
		b.CreateRet(ret)
		b.env = topEnv

		b.SetInsertPointAtEnd(b.entry)
		return v, nil
	}
	panic("unreachable")
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

func (b *Builder) gen(expr ast.Expr, referredFrom string) (llvm.Value, error) {
	switch x := expr.(type) {
	case *ast.Ident:
		if x.Name.Lit == referredFrom {
			return llvm.Value{}, fmt.Errorf("self-reference: %s", x.Name)
		}
		// Note that there is possibility of duplication.
		b.refers[referredFrom] = append(b.refers[referredFrom], x.Name.Lit)
		err := b.checkCR(x.Name.Lit, referredFrom)
		if err != nil {
			return llvm.Value{}, err
		}
		t, found := b.env[x.Name.Lit]
		if !found {
			t, err = b.resolve(x.Name.Lit)
			if err != nil {
				return llvm.Value{}, err
			}
		}
		return t, nil
	case *ast.App:
		var err error
		t, found := b.env[x.FnName.Lit]
		if !found {
			t, err = b.resolve(x.FnName.Lit)
			if err != nil {
				return llvm.Value{}, err
			}
		}
		args := make([]llvm.Value, len(x.Args))
		for i, arg := range x.Args {
			args[i], err = b.gen(arg, referredFrom)
			if err != nil {
				return llvm.Value{}, err
			}
		}
		return b.CreateCall(t, args, "call"), nil
	case *ast.Int:
		n, err := strconv.Atoi(x.X.Lit)
		if err != nil {
			return llvm.Value{}, err
		}
		return llvm.ConstInt(llvm.Int32Type(), uint64(n), false), nil
	case *ast.Add:
		v1, err := b.gen(x.X, referredFrom)
		if err != nil {
			return llvm.Value{}, err
		}
		v2, err := b.gen(x.Y, referredFrom)
		if err != nil {
			return llvm.Value{}, err
		}
		return b.CreateAdd(v1, v2, "add"), nil
	case *ast.Sub:
		v1, err := b.gen(x.X, referredFrom)
		if err != nil {
			return llvm.Value{}, err
		}
		v2, err := b.gen(x.Y, referredFrom)
		if err != nil {
			return llvm.Value{}, err
		}
		return b.CreateSub(v1, v2, "sub"), nil
	case *ast.Mul:
		v1, err := b.gen(x.X, referredFrom)
		if err != nil {
			return llvm.Value{}, err
		}
		v2, err := b.gen(x.Y, referredFrom)
		if err != nil {
			return llvm.Value{}, err
		}
		return b.CreateMul(v1, v2, "mul"), nil
	case *ast.Div:
		v1, err := b.gen(x.X, referredFrom)
		if err != nil {
			return llvm.Value{}, err
		}
		v2, err := b.gen(x.Y, referredFrom)
		if err != nil {
			return llvm.Value{}, err
		}
		return b.CreateUDiv(v1, v2, "div"), nil
	}
	panic("unreachable")
}
