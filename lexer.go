package main

//go:generate goyacc -o parser.go parser.y

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/elpinal/gec/ast"
)

const eof = 0

type exprLexer struct {
	line []byte
	peek rune
	err  error

	expr ast.Expr

	off int // information for error messages
}

func (x *exprLexer) Lex(yylval *yySymType) int {
	for {
		c := x.next()
		switch c {
		case eof:
			return eof
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return x.num(c, yylval)
		case '=', '+', '-', '*', '/':
			return int(c)
		case ' ':
		default:
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
L:
	for {
		c = x.next()
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			add(&b, c)
		default:
			break L
		}
	}
	if c != eof {
		x.peek = c
	}
	n, err := strconv.Atoi(b.String())
	if err != nil {
		x.err = err
		return eof
	}
	yylval.num = n
	return NUM
}

func (x *exprLexer) next() rune {
	if x.peek != eof {
		r := x.peek
		x.peek = eof
		return r
	}
	if len(x.line) == 0 {
		return eof
	}
	c, size := utf8.DecodeRune(x.line)
	x.line = x.line[size:]
	x.off++
	if c == utf8.RuneError && size == 1 {
		x.err = errors.New("next: invalid utf8")
		return x.next()
	}
	return c
}

func (x *exprLexer) Error(s string) {
	x.err = fmt.Errorf("parse error (offset: %d, peek: %q): %s", x.off, x.next(), s)
}

func parse(line []byte) (ast.Expr, error) {
	l := exprLexer{line: line}
	yyParse(&l)
	if l.err != nil {
		return nil, l.err
	}
	return l.expr, nil
}
