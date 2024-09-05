package piscine

func PodiumPosition(podium [][]string) [][]string {
	for j := 1; j < len(podium); j++ {
		for i := 0; i < len(podium)-j; i++ {
			if podium[i][0] > podium[i+1][0] {
				min := podium[i+1]
				podium[i+1] = podium[i]
				podium[i] = min
			}
		}
	}
	return podium
}
