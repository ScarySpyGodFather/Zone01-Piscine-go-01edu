package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(c rune) bool {
	if c == 'a' || c == 'i' || c == 'u' || c == 'e' || c == 'o' || c == 'A' || c == 'I' || c == 'U' || c == 'E' || c == 'O' {
		return true
	}
	return false
}

func printStr(str string) {
	for _, e := range str {
		z01.PrintRune(rune(e))
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args[1:]) >= 1 {
		runes := []rune{}
		for index, arg := range os.Args[1:] {
			for _, char := range arg {
				runes = append(runes, char)
			}
			if index != len(os.Args[1:])-1 {
				runes = append(runes, ' ')
			}
		}
		lastRune := len(runes)
		for j := 0; j <= len(runes)/2; j++ {
			if isVowel(runes[j]) {
				for i := lastRune - 1; i >= j; i-- {
					if isVowel(runes[i]) {
						runes[j], runes[i] = runes[i], runes[j]
						lastRune = i
						break
					}
				}
			}
		}
		printStr(string(runes))
	} else {
		z01.PrintRune('\n')
	}
}
