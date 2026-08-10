package lexer

import "strings"

type Token struct {
	value string
}

func Lex(input string) []Token {

	tokens := []Token{}

	splitInput := strings.Fields(input)

	for _, value := range splitInput {
		tokens = append(tokens, Token{value})
	}

	return tokens
}
