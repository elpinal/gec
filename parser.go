//line parser.y:2
package main

import __yyfmt__ "fmt"

//line parser.y:3
import (
	"github.com/elpinal/gec/ast"
	"github.com/elpinal/gec/token"
)

//line parser.y:12
type yySymType struct {
	yys   int
	top   *ast.WithDecls
	decl  *ast.Decl
	decls []*ast.Decl
	expr  ast.Expr
	token token.Token
}

const ILLEGAL = 57346
const NUM = 57347
const IDENT = 57348
const RARROW = 57349
const BOOL = 57350
const IF = 57351
const THEN = 57352
const ELSE = 57353

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ILLEGAL",
	"NUM",
	"IDENT",
	"RARROW",
	"BOOL",
	"IF",
	"THEN",
	"ELSE",
	"';'",
	"'='",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'\\\\'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:133

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 50

var yyAct = [...]int{

	5, 12, 2, 21, 22, 11, 9, 18, 13, 19,
	29, 14, 6, 24, 16, 17, 13, 10, 25, 14,
	6, 8, 23, 27, 28, 7, 33, 31, 32, 8,
	34, 35, 36, 24, 24, 16, 17, 37, 16, 17,
	15, 26, 13, 19, 20, 14, 30, 3, 4, 1,
}
var yyPact = [...]int{

	11, -1000, -1000, 28, -1000, 24, 37, -1000, 38, -13,
	9, 37, -1000, -1000, -1000, 11, 37, 37, 0, -1000,
	39, 37, 37, 3, -1000, -1000, -1000, -13, -13, 37,
	37, 37, 37, -1000, 21, 24, 37, 24,
}
var yyPgo = [...]int{

	0, 49, 0, 6, 5, 1, 2, 48, 25, 47,
}
var yyR1 = [...]int{

	0, 1, 1, 9, 9, 8, 6, 6, 6, 2,
	2, 2, 3, 3, 3, 4, 4, 5, 5, 5,
	7,
}
var yyR2 = [...]int{

	0, 1, 3, 3, 1, 3, 1, 1, 6, 1,
	3, 3, 1, 3, 3, 1, 2, 1, 1, 1,
	4,
}
var yyChk = [...]int{

	-1000, -1, -6, -9, -7, -2, 9, -8, 18, -3,
	6, -4, -5, 5, 8, 12, 14, 15, -2, 6,
	6, 16, 17, 13, -5, -6, -8, -3, -3, 10,
	7, -4, -4, -6, -2, -2, 11, -2,
}
var yyDef = [...]int{

	0, -2, 1, 0, 6, 7, 0, 4, 0, 9,
	18, 12, 15, 17, 19, 0, 0, 0, 0, 18,
	0, 0, 0, 0, 16, 2, 3, 10, 11, 0,
	0, 13, 14, 5, 0, 20, 0, 8,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 16, 14, 3, 15, 3, 17, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 12,
	3, 13, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 18,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:31
		{
			yyVAL.top = &ast.WithDecls{Expr: yyDollar[1].expr}
			if l, ok := yylex.(*exprLexer); ok {
				l.expr = yyVAL.top
			}
		}
	case 2:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:38
		{
			yyVAL.top = &ast.WithDecls{Decls: yyDollar[1].decls, Expr: yyDollar[3].expr}
			if l, ok := yylex.(*exprLexer); ok {
				l.expr = yyVAL.top
			}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:47
		{
			yyVAL.decls = append(yyDollar[1].decls, yyDollar[3].decl)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:51
		{
			yyVAL.decls = []*ast.Decl{yyDollar[1].decl}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:57
		{
			yyVAL.decl = &ast.Decl{LHS: yyDollar[1].token, RHS: yyDollar[3].expr}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:63
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:67
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 8:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:71
		{
			yyVAL.expr = &ast.If{Cond: yyDollar[2].expr, E1: yyDollar[4].expr, E2: yyDollar[6].expr}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:77
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:81
		{
			yyVAL.expr = &ast.Add{X: yyDollar[1].expr, Y: yyDollar[3].expr}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:85
		{
			yyVAL.expr = &ast.Sub{X: yyDollar[1].expr, Y: yyDollar[3].expr}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:91
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:95
		{
			yyVAL.expr = &ast.Mul{X: yyDollar[1].expr, Y: yyDollar[3].expr}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:99
		{
			yyVAL.expr = &ast.Div{X: yyDollar[1].expr, Y: yyDollar[3].expr}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:105
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:109
		{
			yyVAL.expr = &ast.App{Fn: yyDollar[1].expr, Arg: yyDollar[2].expr}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:115
		{
			yyVAL.expr = &ast.Int{X: yyDollar[1].token}
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:119
		{
			yyVAL.expr = &ast.Ident{Name: yyDollar[1].token}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:123
		{
			yyVAL.expr = &ast.Bool{X: yyDollar[1].token}
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:129
		{
			yyVAL.expr = &ast.Abs{Param: yyDollar[2].token, Body: yyDollar[4].expr}
		}
	}
	goto yystack /* stack new state and value */
}
