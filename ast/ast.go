package ast

import "github.com/elpinal/gec/token"

type Expr interface {
	expr()
}

type Int struct {
	X token.Token
}

func (x *Int) expr() {}

type Bool struct {
	X token.Token
}

func (x *Bool) expr() {}

type Add struct {
	X, Y Expr
}

func (x *Add) expr() {}

type Sub struct {
	X, Y Expr
}

func (x *Sub) expr() {}

type Mul struct {
	X, Y Expr
}

func (x *Mul) expr() {}

type Div struct {
	X, Y Expr
}

func (x *Div) expr() {}

type Ident struct {
	Name token.Token
}

func (x *Ident) expr() {}

type WithDecls struct {
	Decls []*Decl
	Expr  Expr
}

type Decl struct {
	LHS token.Token
	RHS Expr
}

func (x *Decl) LName() string {
	return x.LHS.Lit
}

func (x *Decl) Pos() token.Position {
	return token.NewPosition(x.LHS.Line, x.LHS.Column)
}

type App struct {
	Fn  Expr
	Arg Expr
}

func (x *App) expr() {}

type Abs struct {
	Param token.Token
	Body  Expr
}

func (x *Abs) expr() {}

type If struct {
	Cond Expr
	E1   Expr
	E2   Expr
}

func (x *If) expr() {}

type CmpOp int

const (
	InvalidCmpOp CmpOp = iota
	Eq
	NE
	LT
	GT
	LE
	GE
)

type Cmp struct {
	Op  CmpOp
	LHS Expr
	RHS Expr
}

func (x *Cmp) expr() {}

type NilList struct{}

func (x *NilList) expr() {}

type ParenExpr struct {
	X Expr
}

func (x *ParenExpr) expr() {}
