package piscine

func JumpOver(str string) string {
	if len(str) < 3 {
		return "\n"
	}
	st := ""
	for i := 2; i < len(str); i += 3 {
		st += string(str[i])
	}
	return st + "\n"
}
