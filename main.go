package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		line, _ := reader.ReadString('\n')

		fmt.Print(line)
	}
}
