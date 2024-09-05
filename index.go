package piscine

func Index(s string, toFind string) int {
	a := []rune(s)
	b := []rune(toFind)
	if len(toFind) == 0 || len(toFind) > len(s) {
		return 0
	} else {
		for i := range s {
			if a[i] == b[0] {

				if len(toFind) == 1 {
					return i
				}
				c := 0
				for k := range toFind {
					if len(s) > i+k {
						if a[i+k] == b[k] {
							c++
						}
						if c == len(toFind) {
							return i
						}
					}
				}

			}
		}
		return -1
	}
}
