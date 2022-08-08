package main

import "fmt"

type Token int

const (
	TokenEOF Token = iota + 1
	TokenEOL
	TokenVAR
	TokenDIGIT
	TokenCONST
	TokenIF
	TokenELSE
	TokenPRINT
	TokenWHILE
	TokenIDENTIFIER
	TokenLEFT
	TokenRIGHT
	TokenPLUS
	TokenMINUS
	TokenMUL
	TokenDIV
	TokenASSIGN
	TokenERROR
	TokenEMPTY
)

var (
	_currentLine   = 0
	lastIdentifier IdentifierEntry
)

func IsUpperCase(character uint8) bool {
	return character > 65 && character < 90
}
func IsLowerCase(character uint8) bool {
	return character > 96 && character < 123
}

func CheckWord(code string, word string, currentChar int) bool {
	currentWordChar := 0
	for code[currentChar] == word[currentWordChar] {
		currentChar += 1
		currentWordChar += 1
		if currentWordChar == len(word) && code[currentChar] == ' ' {
			return true
		}
	}
	return false
}

func LexKeyword(code string, currentChar int) (Token, int) {
	switch code[currentChar] {
	case 'P':
		if CheckWord(code, "PRINT", currentChar) {
			return TokenPRINT, currentChar + 6
		}
		break
	case 'V':
		if CheckWord(code, "VAR", currentChar) {
			return TokenVAR, currentChar + 4
		}
		break
	}

	return TokenERROR, currentChar
}

func WordLength(code string, currentChar int) int {
	length := currentChar
	for IsLowerCase(code[length]) {
		length += 1
	}
	return length - currentChar
}

func NextToken(code string, currentChar int) (Token, int) {
	character := code[currentChar]
	switch {
	case character == ' ':
		return NextToken(code, currentChar+1)

	case IsUpperCase(code[currentChar]): // Keywords - Uppercase
		return LexKeyword(code, currentChar)

	case IsLowerCase(code[currentChar]): // Identifiers - Lowercase
		wordLength := WordLength(code, currentChar)
		currentChar += wordLength
		word := Identifier(code[currentChar-wordLength : currentChar])
		InsertIdentifier(word)
		lastIdentifier = hash(word)
		return TokenIDENTIFIER, currentChar

	// Arithmetic operators
	case character == '=':
		return TokenASSIGN, currentChar + 1
	case character == '+':
		return TokenPLUS, currentChar + 1
	case character == '-':
		return TokenMINUS, currentChar + 1
	case character == '*':
		return TokenMUL, currentChar + 1
	case character == '/':
		return TokenDIV, currentChar + 1
	case character == '(':
		return TokenLEFT, currentChar + 1
	case character == ')':
		return TokenRIGHT, currentChar + 1

	case character > 47 && character < 58: // Digits
		return TokenDIGIT, currentChar + 1

	case character == ';': // End of Line - EOL
		_currentLine += 1
		return TokenEOL, currentChar + 1

	case character == '\000': // End of file - EOF
		return TokenEOF, currentChar

	default: // Unidentified symbol
		return TokenERROR, currentChar
	}
}

func PrintTokens(code string) {
	token, currentChar := TokenEMPTY, 0
	for token != TokenEOL {
		token, currentChar = NextToken(code, currentChar)
		fmt.Printf("<%d> at line %d, char %d\n", token, _currentLine, currentChar)
	}
}
