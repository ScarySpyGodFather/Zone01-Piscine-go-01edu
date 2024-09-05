package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := []string(os.Args)
	increment := 48
	if len(os.Args) > 1 {
		if os.Args[1] == "--upper" {
			increment = 16
		}
		for i := 1; i < len(args); i++ {
			if len(args[i]) == 1 {
				if args[i] >= "0" && args[i] <= "9" {
					z01.PrintRune(rune(args[i][0]) + rune(increment))
				} else {
					z01.PrintRune(' ')
				}
			} else {
				number := (args[i][0]-48)*10 + (args[i][1] - 48)
				if number <= 26 && !(args[i] >= "a" && args[i] <= "z") || (args[i] >= "A" && args[i] <= "Z") {
					z01.PrintRune(rune(number+48) + rune(increment))
				} else if increment == 48 {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
