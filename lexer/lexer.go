package lexer

import (
	"errors"
	"slices"
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

func isPitch(word string) bool {
	if len(word) != 2 && len(word) != 3 {
		return false
	}

	_, exists := music.NaturalPitches[word[0]]
	if !exists {
		return false
	}

	if len(word) == 3 && word[1] != music.Flat && word[1] != music.Sharp {
		return false
	}

	octave := word[len(word)-1]

	if octave > music.MaxOctave || octave < music.MinOctave {
		return false
	}

	return true
}

func lexWord(word string) (Token, error) {

	if isPitch(word) {
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
