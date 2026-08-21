package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/adam-oprchal/riff/lexer"
	"github.com/adam-oprchal/riff/midi"
	"github.com/adam-oprchal/riff/parser"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error: cannot read input")
			return
		}

		tokens, err := lexer.Lex(line)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Tokens in line:", tokens)

		program, err := parser.Parse(tokens)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Parsed program:", program)

		midi.PlayProgram(program)
	}
}
