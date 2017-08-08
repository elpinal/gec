package main

import (
	"fmt"
	"os"

	"llvm.org/llvm/bindings/go/llvm"
)

func main() {
	builder := llvm.NewBuilder()
	mod := llvm.NewModule("gec")

	main := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	llvm.AddFunction(mod, "main", main)
	block := llvm.AddBasicBlock(mod.NamedFunction("main"), "entry")
	builder.SetInsertPoint(block, block.FirstInstruction())

	a := builder.CreateAlloca(llvm.Int32Type(), "a")
	expr, err := parse([]byte("12"))
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
	}
	builder.CreateStore(llvm.ConstInt(llvm.Int32Type(), uint64(expr), false), a)

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
