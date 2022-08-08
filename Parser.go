package main

var (
	currentChar                = 0
	token                      Token
	parseUntilRightParenthesis = false
	parserError                ParserError
)

type ParserError int

const (
	NoError = iota + 1
	KeywordError
	ExpectedVARError
	ExpectedIdentifierError
	ExpectedAssignError
	BadRightError
	BadNumericExpressionError
)

type Number struct {
	value string
}

type Factor struct {
	digit      Number
	expression []Expression
}

type Term struct {
	factor   Factor
	operator Token
	term     []Term
}

type Expression struct {
	term       Term
	operator   Token
	expression []Expression
}

type Declaration struct {
	identifier IdentifierEntry
	extraOp    Token
	expression Expression
}

type Print struct {
	expression Expression
}

type Statement struct {
	declaration Declaration
	print       Print
}

func ParseNumber(stmt string) (Number, ParserError) {
	var number Number
	number.value = ""
	for token == TokenDIGIT { // TODO: Support decimal numbers
		number.value += string(stmt[currentChar])
		token, currentChar = NextToken(stmt, currentChar+1)
	}

	currentChar -= 1

	return number, parserError
}

func ParseFactor(stmt string) (Factor, ParserError) {
	var factor Factor
	token, currentChar = NextToken(stmt, currentChar) // GetIdentifier VAR keyword
	switch token {
	case TokenDIGIT:
		currentChar -= 1
		factor.digit, parserError = ParseNumber(stmt)
	case TokenLEFT:
		parseUntilRightParenthesis = true
		factor.expression[0], parserError = ParseExpression(stmt)
	case TokenRIGHT:
		if !parseUntilRightParenthesis {
			return factor, BadRightError
		}
	default:
		return factor, BadNumericExpressionError
	}

	return factor, parserError
}

func ParseTerm(stmt string) (Term, ParserError) {
	var term Term
	term.factor, parserError = ParseFactor(stmt)

	token, currentChar = NextToken(stmt, currentChar) // Get operator

	if token == TokenMUL || token == TokenDIV { // Check for more terms
		term.operator = token
		term.term = make([]Term, 1)
		term.term[0], parserError = ParseTerm(stmt)
		return term, parserError
	}
	term.operator = TokenEMPTY
	currentChar -= 1
	return term, parserError
}

func ParseExpression(stmt string) (Expression, ParserError) {
	var expression Expression
	expression.term, parserError = ParseTerm(stmt)

	token, currentChar = NextToken(stmt, currentChar) // Get operator

	if token == TokenPLUS || token == TokenMINUS {
		expression.operator = token
		expression.expression = make([]Expression, 1)
		expression.expression[0], parserError = ParseExpression(stmt)
		return expression, parserError
	}

	expression.operator = TokenEMPTY
	currentChar -= 1
	return expression, parserError
}

func ParseDeclaration(stmt string) (Declaration, ParserError) {
	var declaration Declaration

	token, currentChar = NextToken(stmt, currentChar) // Get identifier name
	if token != TokenIDENTIFIER {
		return declaration, ExpectedIdentifierError
	}
	declaration.identifier = lastIdentifier

	token, currentChar = NextToken(stmt, currentChar) // Get assign operator
	if token != TokenASSIGN {
		return declaration, ExpectedAssignError
	}
	declaration.extraOp = TokenASSIGN

	declaration.expression, parserError = ParseExpression(stmt)
	return declaration, parserError
}

func ParsePrint(stmt string) (Print, ParserError) {
	var printTree Print
	printTree.expression, parserError = ParseExpression(stmt)
	return printTree, parserError
}

func ParseStatement(stmt string) (Statement, ParserError) {
	var statement Statement
	for token != TokenEOL {
		token, currentChar = NextToken(stmt, currentChar)
		switch token {
		case TokenVAR:
			statement.declaration, parserError = ParseDeclaration(stmt)
		case TokenPRINT:
			statement.print, parserError = ParsePrint(stmt)
		default:
			return statement, KeywordError
		}
		if parserError != NoError {
			return statement, parserError
		}
	}
	return statement, NoError
}
