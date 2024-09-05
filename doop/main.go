package main

import (
	"os"
)

func isNumber(a string) int {
	sign := 1
	start := 0
	if len(a) == 0 {
		return 0
	}
	if a[0] == '-' {
		sign = -1
		start = 1
	} else if a[0] == '+' {
		start = 1
	}

	num := 0
	for i := start; i < len(a); i++ {
		if a[i] < '0' || a[i] > '9' {
			return 0
		}
		num = num*10 + int(a[i]-'0')
	}

	return num * sign
}

func IntToString(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	sign := ""
	if n < 0 {
		n = -n
		sign = "-"
	}
	for n > 0 {
		m := n % 10
		str := string('0' + m)
		result = str + result
		n = n / 10
	}
	return sign + result
}

func Art(op string, num1, num2 int) string {
	var result int
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			return "No division by 0"
		}
	case "%":
		if num2 != 0 {
			result = num1 % num2
		} else {
			return "No modulo by 0"
		}
	default:
		return ""
	}
	if IntToString(result) == "-" {
		return ""
	}
	return IntToString(result)
}

func Verify(a string) bool {
	if a >= "9223372036854775807" && a >= "-9223372036854775808" && len(a) >= len("9223372036854775807") {
		return false
	}
	if len(a) == 0 {
		return false
	}
	start := 0
	if a[0] == '-' || a[0] == '+' {
		start = 1
	}
	for i := start; i < len(a); i++ {
		if a[i] < '0' || a[i] > '9' {
			return false
		}
	}
	return true
}

func operatorValid(op string) bool {
	validOperators := []string{"+", "-", "*", "/", "%"}
	for _, e := range validOperators {
		if op == e {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) == 4 {
		if Verify(os.Args[1]) && Verify(os.Args[3]) {
			nbr1 := isNumber(os.Args[1])
			nbr2 := isNumber(os.Args[3])
			if operatorValid(os.Args[2]) {
				result := Art(os.Args[2], nbr1, nbr2)
				if result != "" {
					os.Stdout.Write([]byte(result))
					os.Stdout.Write([]byte("\n"))
				}
			}
		}
	}
}
