package token

import "fmt"

type Token struct {
	Kind int
	Lit  string
	Position
}

type Position struct {
	Line   uint
	Column uint
}

func NewPosition(line, column uint) Position {
	return Position{
		Line:   line,
		Column: column,
	}
}

func (p Position) String() string {
	return fmt.Sprintf("%d:%d", p.Line, p.Column)
}
