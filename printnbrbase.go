package piscine

import (
	"github.com/01-edu/z01"
)

func PStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func PrintNbrBase(nbr int, base string) {
	breckloop := true
	str := ""
	le := len(base)
	for i, item1 := range base {
		if item1 == '-' || item1 == '+' {
			breckloop = false
			break
		}
		for _, item2 := range base[i+1:] {
			if item1 == item2 {
				breckloop = false
				break
			}
		}
	}
	if breckloop && le >= 2 {
		if nbr == 9223372036854775807 {
			PStr("9223372036854775807")
		} else if nbr == -9223372036854775808 {
			PStr("-9223372036854775808")
		} else {
			if nbr < 0 {

				nbr *= -1
				z01.PrintRune('-')
			}
			for nbr > 0 {
				mod := nbr % le
				str = string(base[mod]) + str
				nbr = nbr / le
			}
			for _, e := range str {
				z01.PrintRune(rune(e))
			}
		}
	} else {
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
}
