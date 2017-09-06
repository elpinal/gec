//line parser.y:2
package parser

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
	cmpop ast.CmpOp
}

const ILLEGAL = 57346
const NEWLINE = 57347
const NUM = 57348
const IDENT = 57349
const SYMBOL = 57350
const BOOL = 57351
const IF = 57352
const THEN = 57353
const ELSE = 57354
const RARROW = 57355
const EQ = 57356
const NE = 57357
const LE = 57358
const GE = 57359

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ILLEGAL",
	"NEWLINE",
	"NUM",
	"IDENT",
	"SYMBOL",
	"BOOL",
	"IF",
	"THEN",
	"ELSE",
	"RARROW",
	"EQ",
	"NE",
	"LE",
	"GE",
	"'='",
	"'\\\\'",
	"'<'",
	"'>'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'('",
	"')'",
	"'['",
	"']'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:202

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 98

var yyAct = [...]int{

	7, 19, 10, 17, 18, 14, 46, 4, 20, 15,
	16, 21, 11, 60, 27, 41, 42, 63, 12, 40,
	44, 13, 64, 45, 50, 51, 47, 43, 22, 39,
	23, 31, 32, 56, 20, 28, 52, 21, 11, 29,
	55, 4, 53, 54, 48, 57, 58, 13, 59, 26,
	25, 2, 62, 61, 22, 24, 23, 30, 8, 20,
	15, 44, 21, 11, 65, 20, 28, 66, 21, 9,
	5, 6, 13, 1, 3, 0, 0, 0, 49, 22,
	0, 23, 0, 0, 0, 22, 0, 23, 33, 34,
	37, 38, 0, 0, 35, 36, 31, 32,
}
var yyPact = [...]int{

	-1000, -1000, 2, -1000, -1000, -1000, -1000, -1000, 44, -1000,
	-1000, 59, -1000, 32, 74, 11, -9, 19, 59, -1000,
	-1000, -1000, 28, -23, 36, 53, -1000, 13, -1000, 12,
	59, 59, 59, -1000, -1000, -1000, -1000, -1000, -1000, 28,
	26, 59, 59, 59, -1000, -14, -1000, -1000, -1000, 36,
	59, 28, 9, -9, -9, -1000, -1, 19, 19, 59,
	-1000, 10, -1000, 28, 59, -1000, -1000,
}
var yyPgo = [...]int{

	0, 74, 73, 71, 70, 5, 10, 4, 3, 1,
	0, 69, 2, 18, 58, 57, 51, 50,
}
var yyR1 = [...]int{

	0, 2, 16, 16, 1, 1, 4, 3, 17, 14,
	14, 13, 13, 10, 10, 10, 11, 12, 12, 15,
	15, 15, 15, 15, 15, 5, 5, 5, 6, 6,
	6, 8, 8, 7, 7, 9, 9, 9, 9, 9,
}
var yyR2 = [...]int{

	0, 3, 0, 2, 1, 1, 1, 3, 2, 3,
	1, 3, 5, 1, 1, 6, 4, 1, 3, 1,
	1, 1, 1, 1, 1, 1, 3, 3, 1, 3,
	3, 1, 3, 1, 2, 1, 1, 1, 3, 2,
}
var yyChk = [...]int{

	-1000, -2, -16, -1, 5, -4, -3, -10, -14, -11,
	-12, 10, -13, 19, -5, 7, -6, -8, -7, -9,
	6, 9, 26, 28, -16, -17, 5, -12, 7, 7,
	-15, 22, 23, 14, 15, 20, 21, 16, 17, 18,
	8, 24, 25, 8, -9, -10, 29, -10, -13, -16,
	11, 13, -5, -6, -6, -10, 7, -8, -8, -7,
	27, -12, -10, 18, 12, -10, -12,
}
var yyDef = [...]int{

	2, -2, 0, 2, 3, 4, 5, 6, 0, 13,
	14, 0, 10, 0, 17, 36, 25, 28, 31, 33,
	35, 37, 0, 0, 1, 0, 2, 0, 36, 0,
	0, 0, 0, 19, 20, 21, 22, 23, 24, 0,
	0, 0, 0, 0, 34, 0, 39, 7, 9, 8,
	0, 0, 18, 26, 27, 11, 0, 29, 30, 32,
	38, 0, 16, 0, 0, 12, 15,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	26, 27, 24, 22, 3, 23, 3, 25, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	20, 18, 21, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 28, 19, 29,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17,
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
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:49
		{
			yyVAL.top = yyDollar[2].top
			if l, ok := yylex.(*exprLexer); ok {
				l.expr = yyVAL.top
			}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:62
		{
			yyVAL.top = &ast.WithDecls{Expr: yyDollar[1].expr}
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:68
		{
			yyVAL.top = &ast.WithDecls{Decls: yyDollar[1].decls, Expr: yyDollar[3].expr}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:76
		{
			yyVAL.decls = append(yyDollar[1].decls, yyDollar[3].decl)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:80
		{
			yyVAL.decls = []*ast.Decl{yyDollar[1].decl}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:86
		{
			yyVAL.decl = &ast.Decl{LHS: yyDollar[1].token, RHS: yyDollar[3].expr}
		}
	case 12:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:90
		{
			f := &ast.Abs{Param: yyDollar[3].token, Body: yyDollar[5].expr}
			g := &ast.Abs{Param: yyDollar[1].token, Body: f}
			yyVAL.decl = &ast.Decl{LHS: yyDollar[2].token, RHS: g}
		}
	case 15:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:100
		{
			yyVAL.expr = &ast.If{Cond: yyDollar[2].expr, E1: yyDollar[4].expr, E2: yyDollar[6].expr}
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:106
		{
			yyVAL.expr = &ast.Abs{Param: yyDollar[2].token, Body: yyDollar[4].expr}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:113
		{
			yyVAL.expr = &ast.Cmp{Op: yyDollar[2].cmpop, LHS: yyDollar[1].expr, RHS: yyDollar[3].expr}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:119
		{
			yyVAL.cmpop = ast.Eq
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:123
		{
			yyVAL.cmpop = ast.NE
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:127
		{
			yyVAL.cmpop = ast.LT
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:131
		{
			yyVAL.cmpop = ast.GT
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:135
		{
			yyVAL.cmpop = ast.LE
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:139
		{
			yyVAL.cmpop = ast.GE
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:146
		{
			yyVAL.expr = &ast.Add{X: yyDollar[1].expr, Y: yyDollar[3].expr}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:150
		{
			yyVAL.expr = &ast.Sub{X: yyDollar[1].expr, Y: yyDollar[3].expr}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:157
		{
			yyVAL.expr = &ast.Mul{X: yyDollar[1].expr, Y: yyDollar[3].expr}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:161
		{
			yyVAL.expr = &ast.Div{X: yyDollar[1].expr, Y: yyDollar[3].expr}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:168
		{
			a := &ast.App{Fn: &ast.Ident{Name: yyDollar[2].token}, Arg: yyDollar[1].expr}
			yyVAL.expr = &ast.App{Fn: a, Arg: yyDollar[3].expr}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:176
		{
			yyVAL.expr = &ast.App{Fn: yyDollar[1].expr, Arg: yyDollar[2].expr}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:182
		{
			yyVAL.expr = &ast.Int{X: yyDollar[1].token}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:186
		{
			yyVAL.expr = &ast.Ident{Name: yyDollar[1].token}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:190
		{
			yyVAL.expr = &ast.Bool{X: yyDollar[1].token}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:194
		{
			yyVAL.expr = &ast.ParenExpr{X: yyDollar[2].expr}
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:198
		{
			yyVAL.expr = &ast.NilList{}
		}
	}
	goto yystack /* stack new state and value */
}
