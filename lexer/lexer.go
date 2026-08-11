package lexer

import (
	"slices"
	"strings"
)

type TokenCategory string

const (
	Pitch    = "PITCH"
	Duration = "DURATION"
	Tempo    = "TEMPO"
	Other    = "OTHER"
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

	possible_durations := []string{"whole", "half", "quarter", "eighth", "sixteenth"}
	if slices.Contains(possible_durations, word) {
		return Token{Duration, word}
	}

	if word == "tempo" {
		return Token{Tempo, word}
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
