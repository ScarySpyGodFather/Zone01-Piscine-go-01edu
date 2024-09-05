package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func (p point) ToString(value int) string {
	digits := []rune{}
	for value > 0 {
		digits = append(digits, rune('0'+value%10))
		value /= 10
	}
	digits[0], digits[1] = digits[1], digits[0]
	return string(digits)
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	strX := points.ToString(points.x)
	strY := points.ToString(points.y)
	str := "x = " + strX + ", y = " + strY
	for _, r := range str {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
