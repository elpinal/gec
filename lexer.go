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
	src  []byte
	peek rune
	err  error

	expr *ast.WithDecls

	off    int // information for error messages
	line   uint
	column uint
}

func isAlphabet(c rune) bool {
	return 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z'
}

func isNumber(c rune) bool {
	return '0' <= c && c <= '9'
}

func (x *exprLexer) Lex(yylval *yySymType) int {
	for {
		c := x.next()
		switch c {
		case eof:
			return eof
		case '=', '+', '-', '*', '/', ';':
			return int(c)
		case ' ':
		default:
			if isAlphabet(c) {
				return x.ident(c, yylval)
			}
			if isNumber(c) {
				return x.num(c, yylval)
			}
			fmt.Fprintf(os.Stderr, "[offset: %d]: invalid character: %[1]U %[1]q\n", x.off, c)
			return ILLEGAL
		}
	}
}

func (x *exprLexer) num(c rune, yylval *yySymType) int {
	add := func(b *bytes.Buffer, c rune) {
		if _, err := b.WriteRune(c); err != nil {
			x.err = fmt.Errorf("WriteRune: %s", err)
		}
	}
	var b bytes.Buffer
	add(&b, c)
	line := x.line
	column := x.column
	for {
		c = x.next()
		if isNumber(c) {
			add(&b, c)
		} else {
			break
		}
	}
	if c != eof {
		x.peek = c
	}
	yylval.token = token.Token{
		Lit:    b.String(),
		Kind:   NUM,
		Line:   line,
		Column: column,
	}
	return NUM
}

func (x *exprLexer) ident(c rune, yylval *yySymType) int {
	add := func(b *bytes.Buffer, c rune) {
		if _, err := b.WriteRune(c); err != nil {
			x.err = fmt.Errorf("WriteRune: %s", err)
		}
	}
	var b bytes.Buffer
	add(&b, c)
	line := x.line
	column := x.column
	for {
		c = x.next()
		if isAlphabet(c) {
			add(&b, c)
		} else {
			break
		}
	}
	if c != eof {
		x.peek = c
	}
	yylval.token = token.Token{
		Lit:    b.String(),
		Kind:   IDENT,
		Line:   line,
		Column: column,
	}
	return IDENT
}

func (x *exprLexer) next() rune {
	if x.peek != eof {
		r := x.peek
		x.peek = eof
		return r
	}
	if len(x.src) == 0 {
		return eof
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
		return x.next()
	}
	return c
}

func (x *exprLexer) Error(s string) {
	x.err = fmt.Errorf("[offset: %d]: %s", x.off, s)
}

func parse(src []byte) (*ast.WithDecls, error) {
	l := exprLexer{src: src}
	yyErrorVerbose = true
	yyParse(&l)
	if l.err != nil {
		return nil, l.err
	}
	return l.expr, nil
}
