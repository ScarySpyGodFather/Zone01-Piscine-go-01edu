package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	path := []rune(os.Args[0])
	name := ""
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			break
		} else {
			name = string(path[i]) + name
		}
	}

	for _, e := range name {
		z01.PrintRune(rune(e))
	}
	z01.PrintRune('\n')
}
