package main

//go:generate goyacc -o parser.go parser.y

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/elpinal/gec/ast"
	"github.com/elpinal/gec/token"
)

const eof = 0

type exprLexer struct {
	src []byte // source
	ch  rune   // current character
	err error

	// result
	expr *ast.WithDecls

	// information for error messages
	off    uint // start at 0
	line   uint // start at 1
	column uint // start at 1

	// information for current token
	tokLine   uint
	tokColumn uint
}

func newLexer(src []byte) *exprLexer {
	l := &exprLexer{
		src:  src,
		line: 1,
	}
	l.next()
	return l
}

func isAlphabet(c rune) bool {
	return 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z'
}

func isNumber(c rune) bool {
	return '0' <= c && c <= '9'
}

func (x *exprLexer) Lex(yylval *yySymType) int {
	for {
		x.tokLine = x.line
		x.tokColumn = x.column
		c := x.ch
		switch c {
		case eof:
			return eof
		case '-':
			x.next()
			if x.ch == '>' {
				x.next()
				return RARROW
			}
			return int(c)
		case '=', '+', '*', '/', ';', '\\':
			x.next()
			return int(c)
		case ' ', '\n':
			x.next()
		default:
			if isAlphabet(c) {
				return x.ident(yylval)
			}
			if isNumber(c) {
				return x.num(yylval)
			}
			fmt.Fprintf(os.Stderr, "[%d:%d]: invalid character: %[3]U %[3]q\n", x.line, x.column, c)
			return ILLEGAL
		}
	}
}

func (x *exprLexer) num(yylval *yySymType) int {
	return x.takeWhile(NUM, isNumber, yylval)
}

func (x *exprLexer) ident(yylval *yySymType) int {
	return x.takeWhile(IDENT, isAlphabet, yylval)
}

func (x *exprLexer) takeWhile(kind int, f func(rune) bool, yylval *yySymType) int {
	add := func(b *bytes.Buffer, c rune) {
		if _, err := b.WriteRune(c); err != nil {
			x.err = fmt.Errorf("WriteRune: %s", err)
		}
	}
	var b bytes.Buffer
	line := x.line
	column := x.column
	for f(x.ch) {
		add(&b, x.ch)
		x.next()
	}
	yylval.token = token.Token{
		Lit:      b.String(),
		Kind:     kind,
		Position: token.NewPosition(line, column),
	}
	return kind
}

func (x *exprLexer) next() {
	if len(x.src) == 0 {
		x.ch = eof
		return
	}
	c, size := utf8.DecodeRune(x.src)
	x.src = x.src[size:]
	x.off++
	if c == '\n' {
		x.line++
		x.column = 0
	} else {
		x.column++
	}
	if c == utf8.RuneError && size == 1 {
		x.err = errors.New("next: invalid utf8")
		x.next()
		return
	}
	x.ch = c
}

func (x *exprLexer) Error(s string) {
	x.err = fmt.Errorf("[%d:%d]: %s", x.tokLine, x.tokColumn, s)
}

func parse(src []byte) (*ast.WithDecls, error) {
	l := newLexer(src)
	yyErrorVerbose = true
	yyParse(l)
	if l.err != nil {
		return nil, l.err
	}
	return l.expr, nil
}
