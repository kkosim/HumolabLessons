package sportsMan

import "fmt"

const weightCof = 13.4
const heightCof = 4.8
const ageCof = 5.7
const bmrCon = 88.36

func GetCharachteristics() (weight, height, age float64, err error) {
	fmt.Print("Введите свой вес в кг, свой пост в см и свой возраст: ")
	_, err = fmt.Scan(&weight, &height, &age)
	if err != nil {
		fmt.Println("Input errror", err.Error())
		return
	}
	if weight <= 0 || height <= 0 || age <= 0 {
		fmt.Println("Давай без шуток")
	}
	return
}

func CountRequiredCalories(weight, height, age float64) float64 {
	dayLimit := bmrCon + (weightCof * weight) + (heightCof * height) - (ageCof * age)
	fmt.Printf("Вам нужно %.2f ккал в день\n", dayLimit)
	return dayLimit
}

func GetDistAndCalEaten() (sprintDist, eatenCal float64, err error) {
	fmt.Print("Сколько километров вы сегодня пробежали? ")
	_, err = fmt.Scan(&sprintDist)
	if err != nil {
		fmt.Println("Input error: ", err.Error())
		return
	}
	fmt.Print("Сколько ккал вы сегодня нахомячили? ")
	_, err = fmt.Scan(&eatenCal)
	if err != nil {
		fmt.Println("Input error: ", err.Error())
		return
	}
	return
}

func CheckBmr(dayLimit, sprintDist, eatenCal, weight float64) {
	burnCal := dayLimit + sprintDist*weight
	if eatenCal > burnCal+10 {
		fmt.Printf("Жрать надо меньше, или пробежи еще %.2f км", (eatenCal-burnCal)/1000)
	} else if eatenCal < burnCal-10 {
		fmt.Printf("Ты на сушке? Если нет, то можешь еще %.2f ккал покушать", burnCal-eatenCal)
	} else {
		fmt.Println("Ай молодец! Твой фитнес тренер гордиться тобой!")
	}
}
