package ast

type Expr interface {
	expr()
}

type Int struct {
	X int
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
	Name string
}

func (x *Ident) expr() {}

type WithDecls struct {
	Decls []Decl
	Expr  Expr
}

type Decl interface {
	LName() string
}

type Assign struct {
	LHS string
	RHS Expr
}

func (x *Assign) LName() string {
	return x.LHS
}

type DeclFunc struct {
	Name string
	Args []string
	RHS  Expr
}

func (x *DeclFunc) LName() string {
	return x.Name
}

type App struct {
	FnName string
	Args   []Expr
}

func (x *App) expr() {}
