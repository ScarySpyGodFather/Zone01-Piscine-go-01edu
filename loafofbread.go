package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	var result string
	var skip string
	count := 0
	for _, char := range str {
		if count%6 != 5 && char == ' ' {
			continue
		}
		if count%6 == 5 {
			skip += " "
		} else {
			skip += string(char)
		}
		count++
	}
	for i := len(skip) - 1; i >= 0; i-- {
		if skip[i] != ' ' {
			result = skip[:i+1]
			break
		}
	}
	return result + "\n"
}
