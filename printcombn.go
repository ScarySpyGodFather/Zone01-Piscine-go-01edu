package piscine

import (
	"github.com/01-edu/z01"
)

func printCombNRecursive(n int, current string, start int, first *bool) {
	if n == 0 {
		if !*first {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		*first = false
		for _, c := range current {
			z01.PrintRune(c)
		}
		return
	}
	for i := start; i <= 9; i++ {
		printCombNRecursive(n-1, current+string('0'+i), i+1, first)
	}
}

func PrintCombN(n int) {
	first := true
	printCombNRecursive(n, "", 0, &first)
	z01.PrintRune('\n')
}
