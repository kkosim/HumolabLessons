//Программа расчитывает время в нормальном виде с правильными окончаниями
//и показывает сколько прошло времени с начала суток. Данные задаются вначале количеством секунд
//прошедших с начала дня

package main

import "fmt"

func main() {
	var n int
	var suffMinutes, suffProshlo1, suffProshlo2, suffHours, suffSeconds string
	fmt.Print("Сколько секунд прошло с начала суток? ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Input error: ", err)
		return
	}
	//Calculating hours, minutes and seconds
	hours := n / 3600
	minutes := (n / 60) % 60
	seconds := n % 60
	//This section is for suffixes of words "прошло" and "час"
	if hours%10 == 1 && hours%100 != 11 {
		suffProshlo1 = "ё"
		suffProshlo2 = ""
		suffHours = ""
	} else if (hours%100)/10 != 1 && hours%10 > 1 && hours%10 < 5 {
		suffProshlo1 = ""
		suffProshlo2 = "o"
		suffHours = "a"
	} else {
		suffProshlo1 = ""
		suffProshlo2 = "o"
		suffHours = "ов"
	}
	//This section is for suffix of "минута"
	if minutes%10 == 1 && minutes != 11 {
		suffMinutes = "а"
	} else if (minutes%10 > 1 && minutes%10 < 5) && (minutes/10 != 1) {
		suffMinutes = "ы"
	} else {
		suffMinutes = ""
	}
	//This section is for suffix of "секунда"
	if seconds%10 == 1 && seconds != 11 {
		suffSeconds = "а"
	} else if (seconds%10 > 1 && seconds%10 < 5) && (seconds/10 != 1) {
		suffSeconds = "ы"
	} else {
		suffSeconds = ""
	}
	fmt.Printf("С начала суток прош%vл%v %v час%v, %v минут%v и %v секунд%v", suffProshlo1, suffProshlo2, hours, suffHours, minutes, suffMinutes, seconds, suffSeconds)
}
