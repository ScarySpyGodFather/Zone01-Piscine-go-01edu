package piscine

func Power2(nb int, power int) int {
	if power < 0 {
		return 0
	}
	result := 1
	for i := 0; i < power; i++ {
		result *= nb
	}
	return result
}

func AtoiBase(s string, base string) int {
	baseLen := len(base)
	if baseLen < 2 {
		return 0
	}

	for i, item1 := range base {
		if item1 == '-' || item1 == '+' {
			return 0
		}
		for _, item2 := range base[i+1:] {
			if item1 == item2 {
				return 0
			}
		}
	}

	for i := 0; i < baseLen; i++ {
		for j := i + 1; j < baseLen; j++ {
			if base[i] == base[j] {
				return 0
			}
		}
	}

	sum := 0
	for i, char := range s {
		index := -1
		for j, baseChar := range base {
			if char == baseChar {
				index = j
				break
			}
		}
		if index == -1 {
			return 0
		}
		sum += index * Power2(len(base), len(s)-i-1)
	}
	return sum
}
