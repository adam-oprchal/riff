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

func runREPL() {

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

		err = midi.PlayProgram(program)

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func readAndCompileFile(input string, output string) {
	fmt.Println("Compling " + input + " to " + output)
}

var logo = `      _  __  __ 
 _ __(_)/ _|/ _|
| '__| | |_| |_ 
| |  | |  _|  _|
|_|  |_|_| |_|  `

func writeHelpMessage() {

	fmt.Println(logo)

	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()

	fmt.Printf("     %-35s %s\n", "riff <input-file> <output-file>", "compile <input-file> into a <output-file> midi file")
	fmt.Printf("     %-35s %s\n", "riff repl", "run an interactive REPL playing sounds in real time with FluidSynth")

	fmt.Println()
}

func main() {

	if len(os.Args) == 2 && os.Args[1] == "repl" {
		runREPL()
	} else if len(os.Args) == 3 {
		readAndCompileFile(os.Args[1], os.Args[2])
	} else {
		writeHelpMessage()
	}
}
