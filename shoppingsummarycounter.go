package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	bar := make(map[string]int)
	var keys string
	for _, vr := range str {
		if vr == ' ' {
			bar[keys] += 1
			keys = ""
		} else {
			keys += string(vr)
		}
	}
	bar[keys] += 1
	return bar
}
