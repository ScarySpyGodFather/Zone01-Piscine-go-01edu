package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	le := len(os.Args)
	if le == 1 {
		T := make([]byte, 1000)
		n, _ := os.Stdin.Read(T)
		if n > 0 {
			T = T[:n-1]
			for string(T) == "Hello" {
				printString(string(T) + "\n")
			}
			printString(string(T) + "\n")
		}
	} else {
		for _, e := range os.Args[1:] {
			file, err := os.Open(e)
			if err != nil {
				printString("ERROR: ")
				printString(err.Error())
				z01.PrintRune('\n')
				os.Exit(1)
			}
			stat, _ := file.Stat()
			Tab := make([]byte, stat.Size())
			file.Read(Tab)
			printString(string(Tab))
			file.Close()
		}
	}
}
