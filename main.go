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

func run(input []byte, logFile *string) {
	builder := llvm.NewBuilder()
	mod := llvm.NewModule("gec")

	main := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	llvm.AddFunction(mod, "main", main)
	block := llvm.AddBasicBlock(mod.NamedFunction("main"), "entry")
	builder.SetInsertPoint(block, block.FirstInstruction())

	expr, err := parse(input)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
	}
	a := gen(builder, expr)

	builder.CreateRet(a)

	if err := llvm.VerifyModule(mod, llvm.ReturnStatusAction); err != nil {
		fmt.Fprintln(os.Stdout, err)
	}
	if logFile != nil {
		ioutil.WriteFile(*logFile, []byte(mod.String()), 0666)
	}

	engine, err := llvm.NewExecutionEngine(mod)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
	}

	funcResult := engine.RunFunction(mod.NamedFunction("main"), []llvm.GenericValue{})
	fmt.Println(funcResult.Int(false))
}

func gen(builder llvm.Builder, expr ast.Expr) llvm.Value {
	switch x := expr.(type) {
	case *ast.Int:
		a := builder.CreateAlloca(llvm.Int32Type(), "a")
		builder.CreateStore(llvm.ConstInt(llvm.Int32Type(), uint64(x.X), false), a)
		return builder.CreateLoad(a, "a")
	case *ast.Add:
		v1 := gen(builder, x.X)
		v2 := gen(builder, x.Y)
		return builder.CreateAdd(v1, v2, "add")
	case *ast.Sub:
		v1 := gen(builder, x.X)
		v2 := gen(builder, x.Y)
		return builder.CreateSub(v1, v2, "sub")
	case *ast.Mul:
		v1 := gen(builder, x.X)
		v2 := gen(builder, x.Y)
		return builder.CreateMul(v1, v2, "sub")
	case *ast.Div:
		v1 := gen(builder, x.X)
		v2 := gen(builder, x.Y)
		return builder.CreateUDiv(v1, v2, "sub")
	}
	panic("unreachable")
}
