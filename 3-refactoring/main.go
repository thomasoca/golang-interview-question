package main

import (
	"fmt"
	"os"
	"strings"
)

func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracket := strings.Index(str, "(")
		indexClosingBracket := strings.Index(str, ")")
		wordsInBracket := str[indexFirstBracket+1 : indexClosingBracket-1]

		return wordsInBracket
	}
	return ""
}

func main() {
	inputStr := os.Args[1]
	outputStr := findFirstStringInBracket(inputStr)
	fmt.Println(outputStr)
}
