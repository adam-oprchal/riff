package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/adam-oprchal/riff/lexer"
	"github.com/adam-oprchal/riff/midi"
	"github.com/adam-oprchal/riff/parser"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("riff> ")

		line, err := reader.ReadString('\n')

		if err == io.EOF {
			fmt.Println("")
			return
		}

		if err != nil {
			fmt.Println("Error: cannot read input")
			return
		}

		tokens, err := lexer.Lex(line)

		if err != nil {
			fmt.Println(err)
			return
		}

		if len(tokens) == 0 {
			continue
		}

		//fmt.Println("Tokens in line:", tokens)

		program, err := parser.Parse(tokens)

		if err != nil {
			fmt.Println(err)
			return
		}

		//fmt.Println("Parsed program:", program)

		fmt.Println("Playing... 🎶")

		midi.PlayProgram(program)
	}
}
