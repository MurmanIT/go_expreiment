package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	my_string := ""
	arguments := os.Args
	if len(arguments) == 1 {
		my_string = "Please set arguments"
	} else {
		my_string = arguments[1]
	}
	io.WriteString(os.Stdout, strings.Join(arguments, " "))
	io.WriteString(os.Stdout, my_string)
	io.WriteString(os.Stdout, "\n")
}
