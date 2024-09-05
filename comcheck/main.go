package main

import (
	"fmt"
	"os"
)

func main() {
	for _, e := range os.Args[1:] {
		if e == "01" || e == "galaxy" || e == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
