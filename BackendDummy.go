package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func DummyTerm(term Term) int {
	if len(term.term) == 0 {
		val, _ := strconv.Atoi(term.factor.digit.value)
		return val
	}
	val1, _ := strconv.Atoi(term.factor.digit.value)
	val2 := DummyTerm(term.term[0])
	switch term.operator {
	case TokenMUL:
		return val2 * val1
	case TokenDIV:
		return val1 / val2
	}

	return 0
}

func DummyExpression(expression Expression) int {
	if len(expression.expression) == 0 {
		val := DummyTerm(expression.term)
		return val
	}
	val1 := DummyTerm(expression.term)
	val2 := DummyExpression(expression.expression[0])
	switch expression.operator {
	case TokenPLUS:
		return val2 + val1
	case TokenMINUS:
		return val1 - val2
	}

	return 0
}

func DummyStatement(stmt Statement) {
	printTree := stmt.print
	if (!reflect.DeepEqual(printTree, Print{})) {
		fmt.Printf("%d", DummyExpression(printTree.expression))
	}
}
