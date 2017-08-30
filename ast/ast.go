package ast

import "github.com/elpinal/gec/token"

type Expr interface {
	expr()
}

type Int struct {
	X token.Token
}

func (x *Int) expr() {}

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
