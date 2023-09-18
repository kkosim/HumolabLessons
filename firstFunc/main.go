package main

import "HumolabLessons/firstFunc/sportsMan"

func main() {
	w, h, a, err := sportsMan.GetCharacteristics()
	if err != nil {
		return
	}

	dayLimit := sportsMan.CountRequiredCalories(w, h, a)

	ran, eaten, err := sportsMan.GetDistAndCalEaten()
	if err != nil {
		return
	}

	sportsMan.CheckBmr(dayLimit, ran, eaten, w)
}
