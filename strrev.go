package piscine

func StrRev(s string) string {
	str := ""
	for _, e := range s {
		str = string(e) + str
	}
	return str
}
