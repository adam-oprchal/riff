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
		line, _ := reader.ReadString('\n')

		tokens := lexer.Lex(line)

		fmt.Println("Tokens in line:", tokens)
	}
}
