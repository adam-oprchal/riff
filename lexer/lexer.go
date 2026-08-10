package lexer

import "strings"

func Lex(input string) []string {

	return strings.Fields(input)
}
