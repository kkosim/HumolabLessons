package rune

import "fmt"

func TwentyGap(fishText string) {
	runes := []rune(fishText)
	for i := 0; i < len(runes); i++ {
		if i%20 == 0 {
			fmt.Println()
		}
		fmt.Print(string(runes[i]))
	}
}
