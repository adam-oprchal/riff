package lexer

import (
	"slices"
	"strings"
)

type TokenCategory string

const (
	Pitch         = "PITCH"
	Duration      = "DURATION"
	Tempo         = "TEMPO"
	NaturalNumber = "NATURALNUMBER"
	Other         = "OTHER"
)

type Token struct {
	category TokenCategory
	value    string
}

func isNaturalNumber(word string) bool {

	for _, char := range word {
		if char < '0' || char > '9' {
			return false
		}
	}

	return word[0] != '0'
}

func lexWord(word string) Token {

	if len(word) == 2 &&
		('A' <= word[0]) && (word[0] <= 'G') &&
		('0' <= word[1]) && (word[1] <= '9') {

		return Token{Pitch, word}
	}

	possible_durations := []string{"whole", "half", "quarter", "eighth", "sixteenth"}
	if slices.Contains(possible_durations, word) {
		return Token{Duration, word}
	}

	if word == "tempo" {
		return Token{Tempo, word}
	}

	if isNaturalNumber(word) {
		return Token{NaturalNumber, word}
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
