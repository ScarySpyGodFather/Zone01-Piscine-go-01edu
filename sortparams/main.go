package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	table := os.Args
	for j := 1; j < len(table); j++ {
		for i := 1; i < len(table)-j; i++ {
			if table[i] > table[i+1] {
				min := table[i+1]
				table[i+1] = table[i]
				table[i] = min
			}
		}
	}
	for i := 1; i < len(table); i++ {
		for _, e := range table[i] {
			z01.PrintRune(e)
		}
		z01.PrintRune('\n')

	}
}
