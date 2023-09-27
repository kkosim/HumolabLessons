package main

import (
	"HumolabLessons/Leetcode"
	"fmt"
)

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extracandies := 3
	fmt.Println(Leetcode.KidsWithCandies(candies, extracandies))

}

/*fish := "Разнообразный и богатый опыт говорит нам, что новая модель организационной деятельности позволяет выполнить важные задания по разработке прогресса профессионального сообщества. В частности, высококачественный прототип будущего проекта предполагает независимые способы реализации позиций, занимаемых участниками в отношении поставленных задач. Однозначно, ключевые особенности структуры проекта будут разоблачены. Мы вынуждены отталкиваться от того, что высокотехнологичная концепция общественного уклада играет важную роль в формировании первоочередных требований."
rune2.TwentyGap(fish)
*/

/*var nums []int
var n int
fmt.Print("Сколько нечетных чисел от 1 печатать? ")
_, err := fmt.Scan(&n)
if err != nil {
	fmt.Println("input error: ", err)
	return
}
for i := 1; i <= n*2-1; i += 2 {
	nums = append(nums, i)
}
fmt.Println(nums)

*/
/*fmt.Println("Вам какую задачу?")
fmt.Println("1 - If28\n2 - If29\n3 - If30\n4 - For39\n5 - For40")
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
}*/

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
