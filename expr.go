package main

import "github.com/elpinal/types-go"

type ArithBinOp struct {
	Op BinOp
	E1 types.Expr
	E2 types.Expr
}

type BinOp int

const (
	InvalidBinOp = iota
	EAdd
	ESub
	EMul
	EDiv
)

func (a *ArithBinOp) Expr() {}
