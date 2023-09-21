package main

import (
	"HumolabLessons/abramyanFor/For39"
	"HumolabLessons/abramyanFor/For40"
	"HumolabLessons/abramyanIf/If28"
	"HumolabLessons/abramyanIf/If29"
	"HumolabLessons/abramyanIf/If30"
	"fmt"
)

func main() {

	fmt.Println("Вам какую задачу?")
	fmt.Println("If28 - 1\nIf29 - 2\nIf30 - 3\nFor39 - 4\nFor40 - 5\n")
	var problem uint8
	_, err := fmt.Scan(&problem)
	if err != nil {
		fmt.Println("Input error: ", err)
		return
	}
	switch problem {
	case 1:
		If28.IsLeapYear()
	case 2:
		If29.EvenOddPosNeg()
	case 3:
		If30.ThreeDigitStatus()
	case 4:
		For39.ABTimes()
	case 5:
		For40.ABSmooth()
	default:
		fmt.Println("Следуйте указаниям!")
	}

	/*id, err := floorAccess.GetEmployeeID()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	floor, err1 := floorAccess.GetFloor()
	if err1 != nil {
		fmt.Println("error: ", err1)
		return
	}
	floorAccess.CheckId(id, floor)
	*/
}

//var somoni float32
/*var somoni int
fmt.Print("How much somoni you have: ")
_, err := fmt.Scan(&somoni)
if err != nil {
	fmt.Println("Ошибка при чтении из терминала:", err.Error())
	return
}
*/
//Converter in golang
/*integer, err := strconv.Atoi("8")
if err != nil {
	fmt.Println("Ошибка конвертации:", err.Error())
	return
}
fmt.Println(integer)

*/
//Contains in golang
/*
	if strings.Contains("Вот какое-то число", "число") {
		fmt.Println("contains")
	} else {
		fmt.Println("doesn't")
	}
*/
//newString := strings.Replace()
//var rubles = somoni * 8.95
//var dollars = somoni * 0.091
/*var rubles = somoni * 895 / 100
var dollars = somoni * 91 / 1000
fmt.Printf("Your money is as valuable as: %v rubles\n", rubles)
fmt.Printf("Your money is as valuable as: %v dollars\n", dollars)
*/
