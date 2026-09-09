package lexer

import (
	"errors"
	"slices"
	"strings"
)

type TokenCategory string

const (
	Pitch         = "PITCH"
	Duration      = "DURATION"
	Tempo         = "TEMPO"
	NaturalNumber = "NATURALNUMBER"
	OpenChord     = "OPENCHORD"
	CloseChord    = "CLOSECHORD"
)

type Token struct {
	Category TokenCategory
	Value    string
}

func isNaturalNumber(word string) bool {

	for _, char := range word {
		if char < '0' || char > '9' {
			return false
		}
	}

	return word[0] != '0'
}

func lexWord(word string) (Token, error) {

	if len(word) == 2 &&
		('A' <= word[0]) && (word[0] <= 'G') &&
		('0' <= word[1]) && (word[1] <= '9') {

		return Token{Pitch, word}, nil
	}

	possible_durations := []string{"whole", "half", "quarter", "eighth", "sixteenth"}
	if slices.Contains(possible_durations, word) {
		return Token{Duration, word}, nil
	}

	if word == "tempo" {
		return Token{Tempo, word}, nil
	}

	if isNaturalNumber(word) {
		return Token{NaturalNumber, word}, nil
	}

	if word == "[" {
		return Token{OpenChord, word}, nil
	}

	if word == "]" {
		return Token{CloseChord, word}, nil
	}

	return Token{}, errors.New("Error: unknown TokenCategory for \"" + word + "\"")
}

func Lex(input string) ([]Token, error) {

	tokens := []Token{}

	splitInput := strings.Fields(input)

	for _, word := range splitInput {

		newToken, err := lexWord(word)

		if err != nil {
			return []Token{}, err
		}

		tokens = append(tokens, newToken)
	}

	return tokens, nil
}
