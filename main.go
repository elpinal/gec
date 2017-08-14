package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/elpinal/gec/ast"

	"llvm.org/llvm/bindings/go/llvm"
)

func main() {
	logFile := flag.String("log", "", "specify `filename` to output LLVM IR")
	flag.Parse()
	if flag.NArg() < 1 {
		return
	}
	run([]byte(flag.Arg(0)), logFile)
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

func run(input []byte, logFile *string) {
	builder := newBuilder(llvm.NewBuilder())
	mod := llvm.NewModule("gec")
	builder.module = mod

	main := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	llvm.AddFunction(builder.module, "main", main)
	block := llvm.AddBasicBlock(builder.module.NamedFunction("main"), "entry")
	builder.entry = block
	builder.SetInsertPoint(block, block.FirstInstruction())

	decls, err := parse(input)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	expr := builder.reserve(decls)
	a := builder.gen(expr, "")

	builder.CreateRet(a)

	if err := llvm.VerifyModule(builder.module, llvm.ReturnStatusAction); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	if logFile != nil {
		ioutil.WriteFile(*logFile, []byte(builder.module.String()), 0666)
	}

	engine, err := llvm.NewExecutionEngine(builder.module)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	funcResult := engine.RunFunction(builder.module.NamedFunction("main"), []llvm.GenericValue{})
	fmt.Println(funcResult.Int(false))
}

func (b *Builder) reserve(wd *ast.WithDecls) ast.Expr {
	for _, decl := range wd.Decls {
		if _, found := b.decls[decl.LName()]; found {
			//TODO: Show previously declared position.
			fmt.Fprintln(os.Stdout, "redeclared:", decl.LName())
			continue
		}
		b.decls[decl.LName()] = decl
	}
	return wd.Expr
}

func (b *Builder) resolve(name string) llvm.Value {
	decl, found := b.decls[name]
	if !found {
		panic(fmt.Sprintf("unknown name: %s", name))
	}
	t := b.genDecl(decl)
	b.env[name] = t
	return t
}

func (b *Builder) genDecl(decl ast.Decl) llvm.Value {
	switch x := decl.(type) {
	case *ast.Assign:
		v := b.gen(x.RHS, x.LHS)
		return v
	case *ast.DeclFunc:
		params := make([]llvm.Type, len(x.Args))
		for i := range x.Args {
			params[i] = llvm.Int32Type()
		}
		f := llvm.FunctionType(llvm.Int32Type(), params, false)
		v := llvm.AddFunction(b.module, x.Name, f)
		for i, name := range x.Args {
			v.Param(i).SetName(name)
		}
		block := llvm.AddBasicBlock(v, "entry")
		b.SetInsertPointAtEnd(block)

		topEnv := make(map[string]llvm.Value, len(b.env))
		for k, v := range b.env {
			topEnv[k] = v
		}
		for i, name := range x.Args {
			b.env[name] = v.Param(i)
		}
		ret := b.gen(x.RHS, x.Name)
		b.CreateRet(ret)
		b.env = topEnv

		b.SetInsertPointAtEnd(b.entry)
		return v
	}
	panic("unreachable")
}

func (b *Builder) checkCR(name, referredFrom string) {
	for _, r := range b.refers[name] {
		if r == referredFrom {
			panic(fmt.Sprintf("circular reference: %s", r))
		}
		b.checkCR(r, referredFrom)
	}
}

func (b *Builder) gen(expr ast.Expr, referredFrom string) llvm.Value {
	switch x := expr.(type) {
	case *ast.Ident:
		if x.Name == referredFrom {
			panic(fmt.Sprintf("self-reference: %s", x.Name))
		}
		// Note that there is possibility of duplication.
		b.refers[referredFrom] = append(b.refers[referredFrom], x.Name)
		b.checkCR(x.Name, referredFrom)
		t, found := b.env[x.Name]
		if !found {
			t = b.resolve(x.Name)
		}
		return t
	case *ast.App:
		t, found := b.env[x.FnName]
		if !found {
			t = b.resolve(x.FnName)
		}
		args := make([]llvm.Value, len(x.Args))
		for i, arg := range x.Args {
			args[i] = b.gen(arg, referredFrom)
		}
		return b.CreateCall(t, args, "call")
	case *ast.Int:
		a := b.CreateAlloca(llvm.Int32Type(), "a")
		b.CreateStore(llvm.ConstInt(llvm.Int32Type(), uint64(x.X), false), a)
		return b.CreateLoad(a, "a")
	case *ast.Add:
		v1 := b.gen(x.X, referredFrom)
		v2 := b.gen(x.Y, referredFrom)
		return b.CreateAdd(v1, v2, "add")
	case *ast.Sub:
		v1 := b.gen(x.X, referredFrom)
		v2 := b.gen(x.Y, referredFrom)
		return b.CreateSub(v1, v2, "sub")
	case *ast.Mul:
		v1 := b.gen(x.X, referredFrom)
		v2 := b.gen(x.Y, referredFrom)
		return b.CreateMul(v1, v2, "mul")
	case *ast.Div:
		v1 := b.gen(x.X, referredFrom)
		v2 := b.gen(x.Y, referredFrom)
		return b.CreateUDiv(v1, v2, "div")
	}
	panic("unreachable")
}
