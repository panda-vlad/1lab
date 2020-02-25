package main

import (
	"errors"
	"fmt"
	lab1 "github.com/Vlad1slavIP74/1lab"
	"log"
	"os"
	s "strings"
)

func main() {
	helper := "usage: git [--help] [--postfix <expression>]"

	if len(os.Args) < 2 {
		err := errors.New("entered incorrect data")
		log.Fatal(err)
	}

	if os.Args[1] == "postfix" {
		res, _ := lab1.PostfixToInfix(s.Join(os.Args[2:], " "))
		fmt.Println(res)
	} else if os.Args[1] == "help" {
		fmt.Println(helper)
	} else {
		fmt.Println("expected 'postfix' or 'help' subcommands")
		os.Exit(1)
	}
}
