package main

import (
	"fmt"
	"os"
)

func help() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")

	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Print("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func sort(table []rune) {
	for j := 1; j < len(table); j++ {
		for i := 0; i < len(table)-j; i++ {
			if table[i] > table[i+1] {
				min := table[i+1]
				table[i+1] = table[i]
				table[i] = min
			}
		}
	}
}

func main() {
	var result string
	args := os.Args[1:]
	if len(os.Args) <= 1 {
		help()
	}
	for _, arg := range args {
		if len(arg) == 0 {
			break
		}
		if arg[0] != '-' {
			result += arg
		}
	}
	for _, arg := range args {
		if len(arg) == 0 {
			break
		}
		if arg[0] == '-' {
			if arg[1] == 'i' && arg[2] == '=' {
				result = result + arg[3:]
			} else if arg[1] == 'o' && len(arg) == 2 {
				table := []rune(result)
				sort(table)
				result = string(table)
			} else if arg[1] == 'h' && len(arg) == 2 {
				help()
			} else if arg[1] == '-' {
				if len(arg) >= 8 && arg[2:8] == "insert" {
					result = result + arg[9:]
				} else if len(arg) == 7 && arg[2:] == "order" {
					table := []rune(result)
					sort(table)
					result = string(table)
				} else if len(arg) == 6 && arg[2:] == "help" {
					help()
				}
			}
		}
	}
	fmt.Println(string(result))
}
