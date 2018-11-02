package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/Chyroc/base8-bagua"
)

var decode bool
var help bool
var helpText = `NAME
     base8-bagua -- Encode and decode using Base8-Bagua representation

SYNOPSIS
     base8-bagua [-D|-h]
`

func init() {
	flag.BoolVar(&decode, "D", false, "decode")
	flag.BoolVar(&help, "h", false, "help")
	flag.Parse()
}

func main() {
	if help {
		fmt.Fprintln(os.Stdout, helpText)
		os.Exit(0)
	}
	line, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if decode {
		result, err := base8_bagua.Decode(string(line))
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, string(result))
	} else {
		result := base8_bagua.Encode(line)
		fmt.Fprintln(os.Stdout, result)
	}
	os.Exit(0)
}
