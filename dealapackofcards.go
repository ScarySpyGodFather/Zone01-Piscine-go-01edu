package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	for j, e := range deck {
		if j < 3 {
			if j == 0 {
				fmt.Printf("Player 1: ")
			}
			if j != 2 {
				fmt.Printf("%v, ", e)
			} else {
				fmt.Printf("%v\n", e)
			}
		} else if j < 6 {
			if j == 3 {
				fmt.Printf("Player 2: ")
			}
			if j != 5 {
				fmt.Printf("%v, ", e)
			} else {
				fmt.Printf("%v\n", e)
			}
		} else if j < 9 {
			if j == 6 {
				fmt.Printf("Player 3: ")
			}
			if j != 8 {
				fmt.Printf("%v, ", e)
			} else {
				fmt.Printf("%v\n", e)
			}
		} else {
			if j == 9 {
				fmt.Printf("Player 4: ")
			}
			if j != 11 {
				fmt.Printf("%v, ", e)
			} else {
				fmt.Printf("%v\n", e)
			}
		}
	}
}
