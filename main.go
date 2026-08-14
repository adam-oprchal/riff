package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/adam-oprchal/riff/lexer"
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
	}
}
