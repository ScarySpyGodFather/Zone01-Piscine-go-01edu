package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i >= 1; i-- {
		for _, e := range os.Args[i] {
			z01.PrintRune(e)
		}
		z01.PrintRune('\n')

	}
}
