package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	moin := n
	result := ""
	if n == 0 {
		z01.PrintRune('0')
	} else {
		for n != 0 {
			m := n % 10
			if m < 0 {
				m = -m
			}
			str := string('0' + m)
			result = str + result
			n = n / 10
		}
		if moin < 0 {
			z01.PrintRune('-')
		}
		for i := 0; i < len(result); i++ {
			z01.PrintRune(rune(result[i]))
		}
	}
}
