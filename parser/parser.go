// Code generated by goyacc -p YY -o parser/parser.go -v parser/parser.output parser/grammar.y. DO NOT EDIT.

//line parser/grammar.y:1

package parser

import __yyfmt__ "fmt"

//line parser/grammar.y:3

//line parser/grammar.y:6
type YYSymType struct {
	yys    int
	String string
	Ast    Ast
}

const NUMBER = 57346
const IDENTIFIER = 57347
const SEPARATOR = 57348
const ASSIGN = 57349
const LET = 57350
const IF = 57351
const THEN = 57352
const LT = 57353
const LTE = 57354
const GT = 57355
const GTE = 57356
const EQ = 57357
const NE = 57358
const OR = 57359
const AND = 57360
const ELSE = 57361
const WHILE = 57362
const PRINT = 57363
const NO_ELSE = 57364
const UMINUS = 57365

var YYToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"IDENTIFIER",
	"SEPARATOR",
	"ASSIGN",
	"LET",
	"IF",
	"THEN",
	"LT",
	"LTE",
	"GT",
	"GTE",
	"EQ",
	"NE",
	"OR",
	"AND",
	"ELSE",
	"WHILE",
	"PRINT",
	"NO_ELSE",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"UMINUS",
	"'('",
	"')'",
	"'{'",
	"'}'",
}

var YYStatenames = [...]string{}

const YYEofCode = 1
const YYErrCode = 2
const YYInitialStackSize = 16

//line parser/grammar.y:84

//line yacctab:1
var YYExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const YYPrivate = 57344

const YYLast = 174

var YYAct = [...]int{
	65, 3, 74, 22, 24, 23, 25, 26, 27, 28,
	29, 70, 68, 34, 36, 18, 19, 20, 21, 72,
	41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
	51, 52, 64, 63, 40, 53, 18, 19, 20, 21,
	56, 57, 58, 39, 38, 71, 22, 24, 23, 25,
	26, 27, 28, 29, 20, 21, 55, 59, 18, 19,
	20, 21, 33, 32, 62, 67, 31, 69, 30, 37,
	66, 1, 2, 73, 22, 24, 23, 25, 26, 27,
	28, 29, 8, 7, 6, 5, 18, 19, 20, 21,
	4, 0, 61, 22, 24, 23, 25, 26, 27, 28,
	29, 0, 0, 0, 0, 18, 19, 20, 21, 0,
	0, 60, 22, 24, 23, 25, 26, 27, 28, 29,
	0, 0, 0, 0, 18, 19, 20, 21, 17, 0,
	54, 9, 35, 22, 24, 23, 25, 26, 27, 28,
	29, 0, 0, 0, 0, 18, 19, 20, 21, 9,
	10, 12, 0, 13, 15, 11, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 16, 14, 0, 0, 12,
	0, 0, 0, 11,
}

var YYPact = [...]int{
	-1000, 145, -1000, 122, 62, 60, 57, -1000, -1000, -1000,
	55, 127, 127, 64, 16, 15, 6, -1000, 127, 127,
	127, 127, 127, 127, 127, 127, 127, 127, 127, 127,
	-1000, -1000, -1000, 127, 101, -1000, -1000, 49, 127, 127,
	127, 29, 29, -1000, -1000, 13, 13, 13, 13, 13,
	13, 13, 13, -8, -1000, 127, 82, 63, 35, -8,
	-1000, 3, 2, 145, 145, -19, 145, -20, 26, -1000,
	-1000, -11, 145, -29, -1000,
}

var YYPgo = [...]int{
	0, 0, 70, 1, 90, 85, 84, 83, 82, 71,
}

var YYR1 = [...]int{
	0, 9, 9, 2, 2, 2, 2, 2, 2, 1,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 4, 5, 6,
	7, 7, 8,
}

var YYR2 = [...]int{
	0, 0, 2, 2, 2, 2, 2, 1, 1, 2,
	1, 1, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 4, 3, 4,
	7, 11, 7,
}

var YYChk = [...]int{
	-1000, -9, -2, -3, -4, -5, -6, -7, -8, 4,
	5, 28, 24, 8, 21, 9, 20, 6, 23, 24,
	25, 26, 11, 13, 12, 14, 15, 16, 17, 18,
	6, 6, 6, 7, -3, 5, -3, 5, 28, 28,
	28, -3, -3, -3, -3, -3, -3, -3, -3, -3,
	-3, -3, -3, -3, 29, 7, -3, -3, -3, -3,
	29, 29, 29, 30, 30, -1, -2, -1, 31, -1,
	31, 19, 30, -1, 31,
}

var YYDef = [...]int{
	1, -2, 2, 0, 0, 0, 0, 7, 8, 11,
	12, 0, 0, 0, 0, 0, 0, 3, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	4, 5, 6, 0, 0, 12, 26, 0, 0, 0,
	0, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 28, 25, 0, 0, 0, 0, 27,
	29, 0, 0, 0, 0, 0, 10, 0, 30, 9,
	32, 0, 0, 0, 31,
}

var YYTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	28, 29, 25, 23, 3, 24, 3, 26, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 30, 3, 31,
}

var YYTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 27,
}

var YYTok3 = [...]int{
	0,
}

var YYErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	YYDebug        = 0
	YYErrorVerbose = false
)

type YYLexer interface {
	Lex(lval *YYSymType) int
	Error(s string)
}

type YYParser interface {
	Parse(YYLexer) int
	Lookahead() int
}

type YYParserImpl struct {
	lval  YYSymType
	stack [YYInitialStackSize]YYSymType
	char  int
}

func (p *YYParserImpl) Lookahead() int {
	return p.char
}

func YYNewParser() YYParser {
	return &YYParserImpl{}
}

const YYFlag = -1000

func YYTokname(c int) string {
	if c >= 1 && c-1 < len(YYToknames) {
		if YYToknames[c-1] != "" {
			return YYToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func YYStatname(s int) string {
	if s >= 0 && s < len(YYStatenames) {
		if YYStatenames[s] != "" {
			return YYStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func YYErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !YYErrorVerbose {
		return "syntax error"
	}

	for _, e := range YYErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + YYTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := YYPact[state]
	for tok := TOKSTART; tok-1 < len(YYToknames); tok++ {
		if n := base + tok; n >= 0 && n < YYLast && YYChk[YYAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if YYDef[state] == -2 {
		i := 0
		for YYExca[i] != -1 || YYExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; YYExca[i] >= 0; i += 2 {
			tok := YYExca[i]
			if tok < TOKSTART || YYExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if YYExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += YYTokname(tok)
	}
	return res
}

func YYlex1(lex YYLexer, lval *YYSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = YYTok1[0]
		goto out
	}
	if char < len(YYTok1) {
		token = YYTok1[char]
		goto out
	}
	if char >= YYPrivate {
		if char < YYPrivate+len(YYTok2) {
			token = YYTok2[char-YYPrivate]
			goto out
		}
	}
	for i := 0; i < len(YYTok3); i += 2 {
		token = YYTok3[i+0]
		if token == char {
			token = YYTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = YYTok2[1] /* unknown char */
	}
	if YYDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", YYTokname(token), uint(char))
	}
	return char, token
}

func YYParse(YYlex YYLexer) int {
	return YYNewParser().Parse(YYlex)
}

func (YYrcvr *YYParserImpl) Parse(YYlex YYLexer) int {
	var YYn int
	var YYVAL YYSymType
	var YYDollar []YYSymType
	_ = YYDollar // silence set and not used
	YYS := YYrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	YYstate := 0
	YYrcvr.char = -1
	YYtoken := -1 // YYrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		YYstate = -1
		YYrcvr.char = -1
		YYtoken = -1
	}()
	YYp := -1
	goto YYstack

ret0:
	return 0

ret1:
	return 1

YYstack:
	/* put a state and value onto the stack */
	if YYDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", YYTokname(YYtoken), YYStatname(YYstate))
	}

	YYp++
	if YYp >= len(YYS) {
		nyys := make([]YYSymType, len(YYS)*2)
		copy(nyys, YYS)
		YYS = nyys
	}
	YYS[YYp] = YYVAL
	YYS[YYp].yys = YYstate

YYnewstate:
	YYn = YYPact[YYstate]
	if YYn <= YYFlag {
		goto YYdefault /* simple state */
	}
	if YYrcvr.char < 0 {
		YYrcvr.char, YYtoken = YYlex1(YYlex, &YYrcvr.lval)
	}
	YYn += YYtoken
	if YYn < 0 || YYn >= YYLast {
		goto YYdefault
	}
	YYn = YYAct[YYn]
	if YYChk[YYn] == YYtoken { /* valid shift */
		YYrcvr.char = -1
		YYtoken = -1
		YYVAL = YYrcvr.lval
		YYstate = YYn
		if Errflag > 0 {
			Errflag--
		}
		goto YYstack
	}

YYdefault:
	/* default state action */
	YYn = YYDef[YYstate]
	if YYn == -2 {
		if YYrcvr.char < 0 {
			YYrcvr.char, YYtoken = YYlex1(YYlex, &YYrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if YYExca[xi+0] == -1 && YYExca[xi+1] == YYstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			YYn = YYExca[xi+0]
			if YYn < 0 || YYn == YYtoken {
				break
			}
		}
		YYn = YYExca[xi+1]
		if YYn < 0 {
			goto ret0
		}
	}
	if YYn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			YYlex.Error(YYErrorMessage(YYstate, YYtoken))
			Nerrs++
			if YYDebug >= 1 {
				__yyfmt__.Printf("%s", YYStatname(YYstate))
				__yyfmt__.Printf(" saw %s\n", YYTokname(YYtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for YYp >= 0 {
				YYn = YYPact[YYS[YYp].yys] + YYErrCode
				if YYn >= 0 && YYn < YYLast {
					YYstate = YYAct[YYn] /* simulate a shift of "error" */
					if YYChk[YYstate] == YYErrCode {
						goto YYstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if YYDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", YYS[YYp].yys)
				}
				YYp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if YYDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", YYTokname(YYtoken))
			}
			if YYtoken == YYEofCode {
				goto ret1
			}
			YYrcvr.char = -1
			YYtoken = -1
			goto YYnewstate /* try again in the same state */
		}
	}

	/* reduction by production YYn */
	if YYDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", YYn, YYStatname(YYstate))
	}

	YYnt := YYn
	YYpt := YYp
	_ = YYpt // guard against "declared and not used"

	YYp -= YYR2[YYn]
	// YYp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if YYp+1 >= len(YYS) {
		nyys := make([]YYSymType, len(YYS)*2)
		copy(nyys, YYS)
		YYS = nyys
	}
	YYVAL = YYS[YYp+1]

	/* consult goto table to find next state */
	YYn = YYR1[YYn]
	YYg := YYPgo[YYn]
	YYj := YYg + YYS[YYp].yys + 1

	if YYj >= YYLast {
		YYstate = YYAct[YYg]
	} else {
		YYstate = YYAct[YYj]
		if YYChk[YYstate] != -YYn {
			YYstate = YYAct[YYg]
		}
	}
	// dummy call; replaced with literal code
	switch YYnt {

	case 2:
		YYDollar = YYS[YYpt-2 : YYpt+1]
//line parser/grammar.y:28
		{
			YYlex.(*Lexer).evalAst(YYDollar[2].Ast)
			PrintAST(YYDollar[2].Ast, 0)
		}
	case 11:
		YYDollar = YYS[YYpt-1 : YYpt+1]
//line parser/grammar.y:46
		{
			YYVAL.Ast = &Number{YYDollar[1].String}
		}
	case 12:
		YYDollar = YYS[YYpt-1 : YYpt+1]
//line parser/grammar.y:47
		{
			YYVAL.Ast = &Variable{YYDollar[1].String}
		}
	case 13:
		YYDollar = YYS[YYpt-3 : YYpt+1]
//line parser/grammar.y:48
		{
			YYVAL.Ast = &BinaryExpr{Op: "+", Lhs: YYDollar[1].Ast, Rhs: YYDollar[3].Ast}
		}
	case 14:
		YYDollar = YYS[YYpt-3 : YYpt+1]
//line parser/grammar.y:49
		{
			YYVAL.Ast = &BinaryExpr{Op: "-", Lhs: YYDollar[1].Ast, Rhs: YYDollar[3].Ast}
		}
	case 15:
		YYDollar = YYS[YYpt-3 : YYpt+1]
//line parser/grammar.y:50
		{
			YYVAL.Ast = &BinaryExpr{Op: "*", Lhs: YYDollar[1].Ast, Rhs: YYDollar[3].Ast}
		}
	case 16:
		YYDollar = YYS[YYpt-3 : YYpt+1]
//line parser/grammar.y:51
		{
			YYVAL.Ast = &BinaryExpr{Op: "/", Lhs: YYDollar[1].Ast, Rhs: YYDollar[3].Ast}
		}
	case 17:
		YYDollar = YYS[YYpt-3 : YYpt+1]
//line parser/grammar.y:52
		{
			YYVAL.Ast = &BinaryExpr{Op: YYDollar[2].String, Lhs: YYDollar[1].Ast, Rhs: YYDollar[3].Ast}
		}
	case 18:
		YYDollar = YYS[YYpt-3 : YYpt+1]
//line parser/grammar.y:53
		{
			YYVAL.Ast = &BinaryExpr{Op: YYDollar[2].String, Lhs: YYDollar[1].Ast, Rhs: YYDollar[3].Ast}
		}
	case 19:
		YYDollar = YYS[YYpt-3 : YYpt+1]
//line parser/grammar.y:54
		{
			YYVAL.Ast = &BinaryExpr{Op: YYDollar[2].String, Lhs: YYDollar[1].Ast, Rhs: YYDollar[3].Ast}
		}
	case 20:
		YYDollar = YYS[YYpt-3 : YYpt+1]
//line parser/grammar.y:55
		{
			YYVAL.Ast = &BinaryExpr{Op: YYDollar[2].String, Lhs: YYDollar[1].Ast, Rhs: YYDollar[3].Ast}
		}
	case 21:
		YYDollar = YYS[YYpt-3 : YYpt+1]
//line parser/grammar.y:56
		{
			YYVAL.Ast = &BinaryExpr{Op: YYDollar[2].String, Lhs: YYDollar[1].Ast, Rhs: YYDollar[3].Ast}
		}
	case 22:
		YYDollar = YYS[YYpt-3 : YYpt+1]
//line parser/grammar.y:57
		{
			YYVAL.Ast = &BinaryExpr{Op: YYDollar[2].String, Lhs: YYDollar[1].Ast, Rhs: YYDollar[3].Ast}
		}
	case 23:
		YYDollar = YYS[YYpt-3 : YYpt+1]
//line parser/grammar.y:58
		{
			YYVAL.Ast = &BinaryExpr{Op: YYDollar[2].String, Lhs: YYDollar[1].Ast, Rhs: YYDollar[3].Ast}
		}
	case 24:
		YYDollar = YYS[YYpt-3 : YYpt+1]
//line parser/grammar.y:59
		{
			YYVAL.Ast = &BinaryExpr{Op: YYDollar[2].String, Lhs: YYDollar[1].Ast, Rhs: YYDollar[3].Ast}
		}
	case 25:
		YYDollar = YYS[YYpt-3 : YYpt+1]
//line parser/grammar.y:60
		{
			YYVAL.Ast = &ParenExpr{YYDollar[2].Ast}
		}
	case 26:
		YYDollar = YYS[YYpt-2 : YYpt+1]
//line parser/grammar.y:61
		{
			YYVAL.Ast = &UnaryExpr{YYDollar[2].Ast}
		}
	case 27:
		YYDollar = YYS[YYpt-4 : YYpt+1]
//line parser/grammar.y:65
		{
			YYVAL.Ast = &Assignment{Variable: YYDollar[2].String, Expr: YYDollar[4].Ast}
		}
	case 28:
		YYDollar = YYS[YYpt-3 : YYpt+1]
//line parser/grammar.y:69
		{
			YYVAL.Ast = &Reassignment{Variable: YYDollar[1].String, Expr: YYDollar[3].Ast}
		}
	case 29:
		YYDollar = YYS[YYpt-4 : YYpt+1]
//line parser/grammar.y:73
		{
			YYVAL.Ast = &StdPrint{Expr: YYDollar[3].Ast}
		}
	case 30:
		YYDollar = YYS[YYpt-7 : YYpt+1]
//line parser/grammar.y:77
		{
			YYVAL.Ast = &IfStatement{Cond: YYDollar[3].Ast, ThenStmt: YYDollar[6].Ast, ElseStmt: nil}
		}
	case 31:
		YYDollar = YYS[YYpt-11 : YYpt+1]
//line parser/grammar.y:78
		{
			YYVAL.Ast = &IfStatement{Cond: YYDollar[3].Ast, ThenStmt: YYDollar[6].Ast, ElseStmt: YYDollar[10].Ast}
		}
	case 32:
		YYDollar = YYS[YYpt-7 : YYpt+1]
//line parser/grammar.y:82
		{
			YYVAL.Ast = &WhileStatement{Cond: YYDollar[3].Ast, Body: YYDollar[6].Ast}
		}
	}
	goto YYstack /* stack new state and value */
}
