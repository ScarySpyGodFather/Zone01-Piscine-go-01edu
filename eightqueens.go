package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	TabJ := map[int]bool{}
	TabJplusI := map[int]bool{}
	TabJmoinI := map[int]bool{}
	result := [8]int{}
	Solve(1, TabJ, TabJplusI, TabJmoinI, result)
}

func PrintRune(arr [8]int) {
	for _, e := range arr {
		z01.PrintRune(rune(e + '0'))
	}
	z01.PrintRune('\n')
}

func Solve(row int, TabJ map[int]bool, TabJplusI map[int]bool, TabJmoinI map[int]bool, result [8]int) bool {
	if row == 9 {
		PrintRune(result)
		return false
	}
	for col := 1; col < 9; col++ {
		if !TabJ[col] && !TabJplusI[col+row] && !TabJmoinI[col-row] {
			TabJ[col] = true
			TabJplusI[col+row] = true
			TabJmoinI[col-row] = true
			result[row-1] = col
			if Solve(row+1, TabJ, TabJplusI, TabJmoinI, result) {
				return true
			}
			TabJ[col] = false
			TabJplusI[col+row] = false
			TabJmoinI[col-row] = false
		}
	}
	return false
}
