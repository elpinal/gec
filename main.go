package main

import (
	"fmt"
	"os"

	"github.com/elpinal/gec/ast"

	"llvm.org/llvm/bindings/go/llvm"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	run([]byte(os.Args[1]))
}

func run(input []byte) {
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
	v := gen(builder, expr)
	a := builder.CreateAlloca(llvm.Int32Type(), "a")
	builder.CreateStore(v, a)

	aVal := builder.CreateLoad(a, "a_val")
	builder.CreateRet(aVal)

	if err := llvm.VerifyModule(mod, llvm.ReturnStatusAction); err != nil {
		fmt.Fprintln(os.Stdout, err)
	}
	//mod.Dump()

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
		return llvm.ConstInt(llvm.Int32Type(), uint64(x.X), false)
	case *ast.Add:
		v1 := gen(builder, x.X)
		v2 := gen(builder, x.Y)
		return llvm.ConstAdd(v1, v2)
	}
	panic("unreachable")
}
