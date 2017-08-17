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
	Decls []Decl
	Expr  Expr
}

type Decl interface {
	LName() string
	Pos() token.Position
}

type Assign struct {
	LHS token.Token
	RHS Expr
}

func (x *Assign) LName() string {
	return x.LHS.Lit
}

func (x *Assign) Pos() token.Position {
	return token.NewPosition(x.LHS.Line, x.LHS.Column)
}

type DeclFunc struct {
	Name token.Token
	Args []token.Token
	RHS  Expr
}

func (x *DeclFunc) LName() string {
	return x.Name.Lit
}

func (x *DeclFunc) Pos() token.Position {
	return token.NewPosition(x.Name.Line, x.Name.Column)
}

type App struct {
	FnName token.Token
	Args   []Expr
}

func (x *App) expr() {}
