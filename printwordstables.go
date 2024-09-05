package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	str := ConcatParams(a)
	for _, e := range str {
		z01.PrintRune(rune(e))
	}
	z01.PrintRune('\n')
}
