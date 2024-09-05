package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else {

		file, _ := os.Open(os.Args[1])
		stat, _ := file.Stat()
		arr := make([]byte, stat.Size())
		file.Read(arr)
		fmt.Println(string(arr))
		file.Close()
	}
}
