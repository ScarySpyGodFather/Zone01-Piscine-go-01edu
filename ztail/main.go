package main

import (
	"os"
)

func AtoiZ(s string) int {
	num := 0
	sign := 1
	for i, element := range s {
		if i == 0 {
			if element == '+' {
				sign = 1
				continue
			} else if element == '-' {
				sign = -1
				continue
			}
		}
		if element < '0' || element > '9' {
			return 0
		}
		basic := int(element - '0')
		num = num*10 + basic
	}
	return num * sign
}

func main() {
	le := len(os.Args)
	if le == 1 {
		os.Exit(1)
	} else if os.Args[1] == "-c" && os.Args[2] >= "0" && os.Args[2] <= "9" {
		if le == 4 {
			size := AtoiZ(os.Args[2])
			file, err := os.Open(os.Args[3])
			if err != nil {
				os.Stdout.Write([]byte(err.Error()))
				os.Stdout.Write([]byte{'\n'})
				os.Exit(1)
			}
			stat, _ := file.Stat()
			Tab := make([]byte, stat.Size())
			file.Read(Tab)
			if size >= len(Tab) {
				for _, e := range Tab {
					os.Stdout.Write([]byte{e})
				}
			} else {
				for _, e := range Tab[len(Tab)-size:] {
					os.Stdout.Write([]byte{e})
				}
			}
			file.Close()
		} else if le > 4 {
			count := 0
			for i, e := range os.Args[3:] {
				size := AtoiZ(os.Args[2])
				file, err := os.Open(e)
				if err != nil {
					os.Stdout.Write([]byte(err.Error()))
					os.Stdout.Write([]byte{'\n'})
					count++
				} else {
					stat, _ := file.Stat()
					Tab := make([]byte, stat.Size())
					file.Read(Tab)
					file.Close()
					if i != 0 {
						os.Stdout.Write([]byte{'\n'})
					}
					os.Stdout.Write([]byte("==> "))
					os.Stdout.Write([]byte(e))
					os.Stdout.Write([]byte(" <=="))
					os.Stdout.Write([]byte{'\n'})
					if size >= len(Tab) {
						for _, e := range Tab {
							os.Stdout.Write([]byte{e})
						}
					} else {
						for _, e := range Tab[len(Tab)-size:] {
							os.Stdout.Write([]byte{e})
						}
					}
				}
			}
			if count > 0 {
				os.Exit(1)
			}
		}
	}
}
