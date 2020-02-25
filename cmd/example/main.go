package main

import (
	"fmt"
	lab "lab1/practice-1"
	"os"
	s "strings"
)

func main() {
	helper := "usage: git [--help] [--postfix <expression>]"

	if os.Args[1] == "postfix" {
		res, _ := lab.PostfixToInfix(s.Join(os.Args[2:], " "))
		fmt.Println(res)
	} else if os.Args[1] == "help" {
		fmt.Println(helper)
	} else {
		fmt.Println("expected 'postfix' or 'help' subcommands")
		os.Exit(1)
	}
}
