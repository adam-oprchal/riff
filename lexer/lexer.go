package lexer

import "strings"

type TokenCategory string

const (
	Pitch = "PITCH"
	Other = "OTHER"
)

type Token struct {
	category TokenCategory
	value    string
}

func lexWord(word string) Token {

	if len(word) == 2 &&
		('A' <= word[0]) && (word[0] <= 'G') &&
		('0' <= word[1]) && (word[1] <= '9') {

		return Token{Pitch, word}
	}

	return Token{Other, word}
}

func Lex(input string) []Token {

	tokens := []Token{}

	splitInput := strings.Fields(input)

	for _, word := range splitInput {
		tokens = append(tokens, lexWord(word))
	}

	return tokens
}
