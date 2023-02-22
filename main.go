package main

import (
	"fmt"

	"github.com/yadiiiig/tfmod/parser"
	"github.com/yadiiiig/tfmod/utils"
)

func main() {
	file, err := utils.OpenFile("_playground/single/main.tf")
	if err != nil {
		fmt.Println(err)
		return
	}

	tokens, err := parser.Lexer(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()

	err = parser.Parse(tokens, file.Content)
	if err != nil {
		fmt.Println(err)
		return
	}
}
