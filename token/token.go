package token

type Token struct {
	Kind   int
	Lit    string
	Line   uint
	Column uint
}
