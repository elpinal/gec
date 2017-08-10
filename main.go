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
	env map[string]llvm.Value
}

func newBuilder(lb llvm.Builder) *Builder {
	return &Builder{
		Builder: lb,
		env:     make(map[string]llvm.Value),
	}
}

func run(input []byte, logFile *string) {
	builder := newBuilder(llvm.NewBuilder())
	mod := llvm.NewModule("gec")

	main := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	llvm.AddFunction(mod, "main", main)
	block := llvm.AddBasicBlock(mod.NamedFunction("main"), "entry")
	builder.SetInsertPoint(block, block.FirstInstruction())

	expr, err := parse(input)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	a := builder.gen(expr)

	builder.CreateRet(a)

	if err := llvm.VerifyModule(mod, llvm.ReturnStatusAction); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	if logFile != nil {
		ioutil.WriteFile(*logFile, []byte(mod.String()), 0666)
	}

	engine, err := llvm.NewExecutionEngine(mod)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	funcResult := engine.RunFunction(mod.NamedFunction("main"), []llvm.GenericValue{})
	fmt.Println(funcResult.Int(false))
}

func (b *Builder) gen(expr ast.Expr) llvm.Value {
	switch x := expr.(type) {
	case *ast.WithDecls:
		for _, decl := range x.Decls {
			_ = b.gen(decl)
		}
		return b.gen(x.Expr)
	case *ast.Assign:
		t := b.CreateAlloca(llvm.Int32Type(), "assign")
		b.CreateStore(b.gen(x.RHS), t)
		if _, found := b.env[x.LHS]; found {
			//TODO: Show previously declared position.
			fmt.Fprintln(os.Stdout, "redeclared:", x.LHS)
		}
		b.env[x.LHS] = t
		return llvm.Value{}
	case *ast.Int:
		a := b.CreateAlloca(llvm.Int32Type(), "a")
		b.CreateStore(llvm.ConstInt(llvm.Int32Type(), uint64(x.X), false), a)
		return b.CreateLoad(a, "a")
	case *ast.Add:
		v1 := b.gen(x.X)
		v2 := b.gen(x.Y)
		return b.CreateAdd(v1, v2, "add")
	case *ast.Sub:
		v1 := b.gen(x.X)
		v2 := b.gen(x.Y)
		return b.CreateSub(v1, v2, "sub")
	case *ast.Mul:
		v1 := b.gen(x.X)
		v2 := b.gen(x.Y)
		return b.CreateMul(v1, v2, "sub")
	case *ast.Div:
		v1 := b.gen(x.X)
		v2 := b.gen(x.Y)
		return b.CreateUDiv(v1, v2, "sub")
	}
	panic("unreachable")
}
