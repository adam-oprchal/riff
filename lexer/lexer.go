package lexer

import (
	"errors"
	"strings"

	"github.com/adam-oprchal/riff/music"
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

	if music.IsPitch(word) {
		return Token{Pitch, word}, nil
	}

	if music.IsDuration(word) {
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

	return Token{}, errors.New("Error: unknown input \"" + word + "\"")
}

func Lex(input string) ([]Token, error) {

	tokens := []Token{}

	replacer := strings.NewReplacer(
		"[", " [ ",
		"]", " ] ",
	)

	splitInput := strings.Fields(replacer.Replace(input))

	for _, word := range splitInput {

		newToken, err := lexWord(word)

		if err != nil {
			return []Token{}, err
		}

		tokens = append(tokens, newToken)
	}

	return tokens, nil
}
