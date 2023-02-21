package main

import (
	"fmt"

	"github.com/yadiiiig/tfmod/lx"
	"github.com/yadiiiig/tfmod/utils"
)

func main() {
	file, err := utils.OpenFile("_playground/single/main.tf")
	if err != nil {
		fmt.Println(err)
		return
	}

	tokens, err := lx.RunLexer(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()

	err = lx.Parser(tokens, file.Content)
	if err != nil {
		fmt.Println(err)
		return
	}
}
