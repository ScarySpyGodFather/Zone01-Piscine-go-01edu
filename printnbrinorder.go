package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var table []rune
	for n > 0 {
		singleDigit := n % 10
		n /= 10
		table = append(table, rune(singleDigit))
	}
	for j := 1; j < len(table); j++ {
		for i := 0; i < len(table)-j; i++ {
			if table[i] > table[i+1] {
				table[i], table[i+1] = table[i+1], table[i]
			}
		}
	}
	for _, e := range table {
		z01.PrintRune(rune(e + 48))
	}
}
