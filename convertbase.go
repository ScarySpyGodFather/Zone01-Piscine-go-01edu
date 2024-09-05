package piscine

func Power(nb int, power int) int {
	if nb == 0 && power == 0 {
		return 1
	}

	if power == 0 {
		return 1
	}

	if power < 0 {
		return 0
	}
	return nb * Power(nb, power-1)
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	le := len(baseFrom)
	tole := len(baseTo)
	nbrInt := 0
	for i := 0; i < len(nbr); i++ {
		nbrInt += int(nbr[len(nbr)-1-i]-48) * Power(le, i)
	}
	newNbr := ""
	for nbrInt != 0 {
		mod := nbrInt % tole
		nbrInt = nbrInt / tole
		newNbr = string(baseTo[mod]) + newNbr
	}
	return newNbr
}
