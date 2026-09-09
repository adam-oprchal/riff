package parser

import (
	"errors"
	"strconv"

	"github.com/adam-oprchal/riff/lexer"
)

type Pitch struct {
	Letter byte
	Octave int
}

type Duration string

type Event interface {
	isEvent()
}

type Note struct {
	Pitch    Pitch
	Duration Duration
}

func (Note) isEvent() {}

type TempoChange struct {
	Value int
}

func (TempoChange) isEvent() {}

type Chord struct {
	Pitches  []Pitch
	Duration Duration
}

func (Chord) isEvent() {}

type Program struct {
	Events []Event
}

func parseNote(tokens []lexer.Token, events *[]Event, curTokenIndex *int) {

	pitchLetter := tokens[*curTokenIndex].Value[0]
	pitchOctave := tokens[*curTokenIndex].Value[1] - '0'

	newPitch := Pitch{pitchLetter, int(pitchOctave)}

	*curTokenIndex++

	if *curTokenIndex >= len(tokens) || tokens[*curTokenIndex].Category != lexer.Duration {
		*events = append(*events, Note{newPitch, "whole"})
		return
	}

	*events = append(*events, Note{newPitch, Duration(tokens[*curTokenIndex].Value)})

	*curTokenIndex++
}

func parseTempo(tokens []lexer.Token, events *[]Event, curTokenIndex *int) error {

	*curTokenIndex++

	if *curTokenIndex >= len(tokens) || tokens[*curTokenIndex].Category != lexer.NaturalNumber {
		return errors.New("Error: expected NaturalNumber after Tempo Token")
	}

	n, _ := strconv.Atoi(tokens[*curTokenIndex].Value)

	*curTokenIndex++

	*events = append(*events, TempoChange{n})

	return nil
}

func parseChord(tokens []lexer.Token, events *[]Event, curTokenIndex *int) error {

	*curTokenIndex++

	pitches := []Pitch{}

	for *curTokenIndex < len(tokens) && tokens[*curTokenIndex].Category == lexer.Pitch {

		pitchLetter := tokens[*curTokenIndex].Value[0]
		pitchOctave := tokens[*curTokenIndex].Value[1] - '0'

		newPitch := Pitch{pitchLetter, int(pitchOctave)}
		pitches = append(pitches, newPitch)

		*curTokenIndex++
	}

	if *curTokenIndex >= len(tokens) || tokens[*curTokenIndex].Category != lexer.CloseChord {
		return errors.New("Error: expected CloseChord after OpenChord/Pitch")
	}

	*curTokenIndex++

	if *curTokenIndex >= len(tokens) || tokens[*curTokenIndex].Category != lexer.Duration {
		*events = append(*events, Chord{pitches, "whole"})
		return nil
	}

	*events = append(*events, Chord{pitches, Duration(tokens[*curTokenIndex].Value)})

	*curTokenIndex++

	return nil
}

func Parse(tokens []lexer.Token) (Program, error) {

	events := []Event{}
	curTokenIndex := 0

	for curTokenIndex < len(tokens) {

		curToken := tokens[curTokenIndex]

		switch curToken.Category {
		case lexer.Pitch:
			parseNote(tokens, &events, &curTokenIndex)
		case lexer.Tempo:
			err := parseTempo(tokens, &events, &curTokenIndex)
			if err != nil {
				return Program{}, err
			}
		case lexer.OpenChord:
			err := parseChord(tokens, &events, &curTokenIndex)
			if err != nil {
				return Program{}, err
			}
		default:
			return Program{}, errors.New("Error: unexpected Token of value " + curToken.Value)
		}
	}

	return Program{events}, nil
}
