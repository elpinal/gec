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

//line parser.y:194

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 96

var yyAct = [...]int{

	7, 19, 10, 17, 18, 16, 45, 4, 20, 15,
	64, 21, 11, 2, 27, 20, 28, 24, 21, 39,
	43, 13, 14, 44, 40, 41, 46, 50, 22, 38,
	23, 67, 36, 37, 68, 22, 49, 23, 12, 59,
	48, 42, 57, 58, 61, 62, 60, 63, 29, 4,
	26, 66, 65, 51, 52, 53, 54, 55, 56, 20,
	28, 25, 21, 11, 47, 43, 8, 9, 69, 5,
	6, 70, 13, 20, 15, 1, 21, 11, 3, 22,
	0, 23, 30, 31, 34, 35, 13, 0, 32, 33,
	36, 37, 0, 22, 0, 23,
}
var yyPact = [...]int{

	-1000, -1000, 2, -1000, -1000, -1000, -1000, -1000, 45, -1000,
	-1000, 9, -1000, 41, 68, 11, 0, 33, 9, -1000,
	-1000, -1000, 53, -23, 44, 67, -1000, 25, -1000, 14,
	9, 9, 9, 9, 9, 9, 9, 9, 53, 39,
	9, 9, 9, -1000, -17, -1000, -1000, -1000, 44, 9,
	53, 10, 10, 10, 10, 10, 10, 0, 0, -1000,
	13, 33, 33, 9, -1000, 22, -1000, 53, 9, -1000,
	-1000,
}
var yyPgo = [...]int{

	0, 78, 75, 70, 69, 22, 5, 4, 3, 1,
	0, 67, 2, 38, 66, 13, 61,
}
var yyR1 = [...]int{

	0, 2, 15, 15, 1, 1, 4, 3, 16, 14,
	14, 13, 13, 10, 10, 10, 11, 12, 12, 12,
	12, 12, 12, 12, 5, 5, 5, 6, 6, 6,
	8, 8, 7, 7, 9, 9, 9, 9, 9,
}
var yyR2 = [...]int{

	0, 3, 0, 2, 1, 1, 1, 3, 2, 3,
	1, 3, 5, 1, 1, 6, 4, 1, 3, 3,
	3, 3, 3, 3, 1, 3, 3, 1, 3, 3,
	1, 3, 1, 2, 1, 1, 1, 3, 2,
}
var yyChk = [...]int{

	-1000, -2, -15, -1, 5, -4, -3, -10, -14, -11,
	-12, 10, -13, 19, -5, 7, -6, -8, -7, -9,
	6, 9, 26, 28, -15, -16, 5, -12, 7, 7,
	14, 15, 20, 21, 16, 17, 22, 23, 18, 8,
	24, 25, 8, -9, -10, 29, -10, -13, -15, 11,
	13, -5, -5, -5, -5, -5, -5, -6, -6, -10,
	7, -8, -8, -7, 27, -12, -10, 18, 12, -10,
	-12,
}
var yyDef = [...]int{

	2, -2, 0, 2, 3, 4, 5, 6, 0, 13,
	14, 0, 10, 0, 17, 35, 24, 27, 30, 32,
	34, 36, 0, 0, 1, 0, 2, 0, 35, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 33, 0, 38, 7, 9, 8, 0,
	0, 18, 19, 20, 21, 22, 23, 25, 26, 11,
	0, 28, 29, 31, 37, 0, 16, 0, 0, 12,
	15,
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
		//line parser.y:47
		{
			yyVAL.top = yyDollar[2].top
			if l, ok := yylex.(*exprLexer); ok {
				l.expr = yyVAL.top
			}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:60
		{
			yyVAL.top = &ast.WithDecls{Expr: yyDollar[1].expr}
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:66
		{
			yyVAL.top = &ast.WithDecls{Decls: yyDollar[1].decls, Expr: yyDollar[3].expr}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:74
		{
			yyVAL.decls = append(yyDollar[1].decls, yyDollar[3].decl)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:78
		{
			yyVAL.decls = []*ast.Decl{yyDollar[1].decl}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:84
		{
			yyVAL.decl = &ast.Decl{LHS: yyDollar[1].token, RHS: yyDollar[3].expr}
		}
	case 12:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:88
		{
			f := &ast.Abs{Param: yyDollar[3].token, Body: yyDollar[5].expr}
			g := &ast.Abs{Param: yyDollar[1].token, Body: f}
			yyVAL.decl = &ast.Decl{LHS: yyDollar[2].token, RHS: g}
		}
	case 15:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:98
		{
			yyVAL.expr = &ast.If{Cond: yyDollar[2].expr, E1: yyDollar[4].expr, E2: yyDollar[6].expr}
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:104
		{
			yyVAL.expr = &ast.Abs{Param: yyDollar[2].token, Body: yyDollar[4].expr}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:111
		{
			yyVAL.expr = &ast.Cmp{Op: ast.Eq, LHS: yyDollar[1].expr, RHS: yyDollar[3].expr}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:115
		{
			yyVAL.expr = &ast.Cmp{Op: ast.NE, LHS: yyDollar[1].expr, RHS: yyDollar[3].expr}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:119
		{
			yyVAL.expr = &ast.Cmp{Op: ast.LT, LHS: yyDollar[1].expr, RHS: yyDollar[3].expr}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:123
		{
			yyVAL.expr = &ast.Cmp{Op: ast.GT, LHS: yyDollar[1].expr, RHS: yyDollar[3].expr}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:127
		{
			yyVAL.expr = &ast.Cmp{Op: ast.LE, LHS: yyDollar[1].expr, RHS: yyDollar[3].expr}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:131
		{
			yyVAL.expr = &ast.Cmp{Op: ast.GE, LHS: yyDollar[1].expr, RHS: yyDollar[3].expr}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:138
		{
			yyVAL.expr = &ast.Add{X: yyDollar[1].expr, Y: yyDollar[3].expr}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:142
		{
			yyVAL.expr = &ast.Sub{X: yyDollar[1].expr, Y: yyDollar[3].expr}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:149
		{
			yyVAL.expr = &ast.Mul{X: yyDollar[1].expr, Y: yyDollar[3].expr}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:153
		{
			yyVAL.expr = &ast.Div{X: yyDollar[1].expr, Y: yyDollar[3].expr}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:160
		{
			a := &ast.App{Fn: yyDollar[2].token, Arg: yyDollar[1].expr}
			yyVAL.expr = &ast.App{Fn: a, Arg: yyDollar[3].expr}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:168
		{
			yyVAL.expr = &ast.App{Fn: yyDollar[1].expr, Arg: yyDollar[2].expr}
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:174
		{
			yyVAL.expr = &ast.Int{X: yyDollar[1].token}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:178
		{
			yyVAL.expr = &ast.Ident{Name: yyDollar[1].token}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:182
		{
			yyVAL.expr = &ast.Bool{X: yyDollar[1].token}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:186
		{
			yyVAL.expr = &ast.ParenExpr{X: yyDollar[2].expr}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:190
		{
			yyVAL.expr = &ast.NilList{}
		}
	}
	goto yystack /* stack new state and value */
}
