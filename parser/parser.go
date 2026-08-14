package parser

import (
	"github.com/adam-oprchal/riff/lexer"
)

type Event interface {
}

type Program struct {
	events []Event
}

func Parse(tokens []lexer.Token) (Program, error) {

	return Program{}, nil
}
