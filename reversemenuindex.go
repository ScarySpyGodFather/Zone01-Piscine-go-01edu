package piscine

func ReverseMenuIndex(menu []string) []string {
	rev := make([]string, len(menu))

	for i := len(menu) - 1; i >= 0; i-- {
		rev[len(menu)-i-1] = menu[i]
	}
	return rev
}
