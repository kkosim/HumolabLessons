package main

import "firstClass/firstFunc/sportsMan"

func main() {
	w, h, a, err := sportsMan.GetCharachteristics()
	if err != nil {
		return
	}
	dayLimit := sportsMan.CountRequiredCalories(w, h, a)
	runned, eaten, err := sportsMan.GetDistAndCalEaten()
	if err != nil {
		return
	}

	sportsMan.CheckBmr(dayLimit, runned, eaten, w)
}
