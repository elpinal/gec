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
const RARROW = 57350
const BOOL = 57351
const IF = 57352
const THEN = 57353
const ELSE = 57354
const EQ = 57355
const NE = 57356
const LE = 57357
const GE = 57358
const SYMBOL = 57359

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ILLEGAL",
	"NEWLINE",
	"NUM",
	"IDENT",
	"RARROW",
	"BOOL",
	"IF",
	"THEN",
	"ELSE",
	"EQ",
	"NE",
	"LE",
	"GE",
	"SYMBOL",
	"'='",
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
	"'\\\\'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:184

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 101

var yyAct = [...]int{

	5, 16, 8, 15, 4, 17, 13, 41, 18, 9,
	14, 59, 24, 17, 25, 10, 18, 39, 37, 38,
	40, 12, 62, 42, 19, 63, 20, 45, 11, 33,
	34, 2, 19, 46, 20, 21, 55, 56, 43, 36,
	35, 57, 58, 26, 53, 54, 4, 61, 60, 47,
	48, 49, 50, 51, 52, 44, 23, 22, 6, 39,
	39, 17, 25, 64, 18, 9, 65, 17, 13, 7,
	18, 9, 1, 3, 0, 0, 0, 0, 0, 0,
	19, 0, 20, 0, 11, 0, 19, 0, 20, 0,
	11, 27, 28, 31, 32, 0, 0, 29, 30, 33,
	34,
}
var yyPact = [...]int{

	-1000, -1000, -1, -1000, -1000, -1000, 51, -1000, -1000, 7,
	-1000, 36, 78, 22, -5, 7, -1000, -1000, -1000, 55,
	-21, 41, 61, -1000, 16, -1000, 25, 7, 7, 7,
	7, 7, 7, 7, 7, 55, 30, 7, 7, -1000,
	-15, -1000, -1000, -1000, 41, 7, 55, 8, 8, 8,
	8, 8, 8, -5, -5, -1000, 4, 7, 7, -1000,
	13, -1000, 55, 7, -1000, -1000,
}
var yyPgo = [...]int{

	0, 73, 72, 21, 10, 3, 1, 0, 69, 2,
	15, 58, 31, 57,
}
var yyR1 = [...]int{

	0, 2, 1, 1, 12, 12, 13, 11, 11, 10,
	10, 7, 7, 7, 9, 9, 9, 9, 9, 9,
	9, 3, 3, 3, 4, 4, 4, 5, 5, 6,
	6, 6, 6, 6, 8,
}
var yyR2 = [...]int{

	0, 3, 1, 3, 0, 2, 2, 3, 1, 3,
	5, 1, 1, 6, 1, 3, 3, 3, 3, 3,
	3, 1, 3, 3, 1, 3, 3, 1, 2, 1,
	1, 1, 3, 2, 4,
}
var yyChk = [...]int{

	-1000, -2, -12, -1, 5, -7, -11, -8, -9, 10,
	-10, 29, -3, 7, -4, -5, -6, 6, 9, 25,
	27, -12, -13, 5, -9, 7, 7, 13, 14, 19,
	20, 15, 16, 21, 22, 18, 17, 23, 24, -6,
	-7, 28, -7, -10, -12, 11, 8, -3, -3, -3,
	-3, -3, -3, -4, -4, -7, 7, -5, -5, 26,
	-9, -7, 18, 12, -7, -9,
}
var yyDef = [...]int{

	4, -2, 0, 4, 5, 2, 0, 11, 12, 0,
	8, 0, 14, 30, 21, 24, 27, 29, 31, 0,
	0, 1, 0, 4, 0, 30, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 28,
	0, 33, 3, 7, 6, 0, 0, 15, 16, 17,
	18, 19, 20, 22, 23, 9, 0, 25, 26, 32,
	0, 34, 0, 0, 10, 13,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	25, 26, 23, 21, 3, 22, 3, 24, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	19, 18, 20, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 27, 29, 28,
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
		//line parser.y:31
		{
			yyVAL.top = yyDollar[2].top
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:37
		{
			yyVAL.top = &ast.WithDecls{Expr: yyDollar[1].expr}
			if l, ok := yylex.(*exprLexer); ok {
				l.expr = yyVAL.top
			}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:44
		{
			yyVAL.top = &ast.WithDecls{Decls: yyDollar[1].decls, Expr: yyDollar[3].expr}
			if l, ok := yylex.(*exprLexer); ok {
				l.expr = yyVAL.top
			}
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:57
		{
			yyVAL.decls = append(yyDollar[1].decls, yyDollar[3].decl)
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:61
		{
			yyVAL.decls = []*ast.Decl{yyDollar[1].decl}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:67
		{
			yyVAL.decl = &ast.Decl{LHS: yyDollar[1].token, RHS: yyDollar[3].expr}
		}
	case 10:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:71
		{
			f := &ast.Abs{Param: yyDollar[3].token, Body: yyDollar[5].expr}
			g := &ast.Abs{Param: yyDollar[1].token, Body: f}
			yyVAL.decl = &ast.Decl{LHS: yyDollar[2].token, RHS: g}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:79
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:83
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 13:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:87
		{
			yyVAL.expr = &ast.If{Cond: yyDollar[2].expr, E1: yyDollar[4].expr, E2: yyDollar[6].expr}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:94
		{
			yyVAL.expr = &ast.Cmp{Op: ast.Eq, LHS: yyDollar[1].expr, RHS: yyDollar[3].expr}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:98
		{
			yyVAL.expr = &ast.Cmp{Op: ast.NE, LHS: yyDollar[1].expr, RHS: yyDollar[3].expr}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:102
		{
			yyVAL.expr = &ast.Cmp{Op: ast.LT, LHS: yyDollar[1].expr, RHS: yyDollar[3].expr}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:106
		{
			yyVAL.expr = &ast.Cmp{Op: ast.GT, LHS: yyDollar[1].expr, RHS: yyDollar[3].expr}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:110
		{
			yyVAL.expr = &ast.Cmp{Op: ast.LE, LHS: yyDollar[1].expr, RHS: yyDollar[3].expr}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:114
		{
			yyVAL.expr = &ast.Cmp{Op: ast.GE, LHS: yyDollar[1].expr, RHS: yyDollar[3].expr}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:120
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:124
		{
			yyVAL.expr = &ast.Add{X: yyDollar[1].expr, Y: yyDollar[3].expr}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:128
		{
			yyVAL.expr = &ast.Sub{X: yyDollar[1].expr, Y: yyDollar[3].expr}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:134
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:138
		{
			yyVAL.expr = &ast.Mul{X: yyDollar[1].expr, Y: yyDollar[3].expr}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:142
		{
			yyVAL.expr = &ast.Div{X: yyDollar[1].expr, Y: yyDollar[3].expr}
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:148
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:152
		{
			yyVAL.expr = &ast.App{Fn: yyDollar[1].expr, Arg: yyDollar[2].expr}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:158
		{
			yyVAL.expr = &ast.Int{X: yyDollar[1].token}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:162
		{
			yyVAL.expr = &ast.Ident{Name: yyDollar[1].token}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:166
		{
			yyVAL.expr = &ast.Bool{X: yyDollar[1].token}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:170
		{
			yyVAL.expr = &ast.ParenExpr{X: yyDollar[2].expr}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:174
		{
			yyVAL.expr = &ast.NilList{}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:180
		{
			yyVAL.expr = &ast.Abs{Param: yyDollar[2].token, Body: yyDollar[4].expr}
		}
	}
	goto yystack /* stack new state and value */
}
