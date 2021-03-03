package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "Trim newline character from the end")
var sep = flag.String("sep", " ", "Arguments separator character")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}
