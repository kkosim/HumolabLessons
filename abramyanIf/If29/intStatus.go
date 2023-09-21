package If29

import "fmt"

func EvenOddPosNeg() {
	fmt.Println("Дано целое число. Вывести его строку-описание вида «отрицательное\nчетное число», «нулевое число», «положительное нечетное число» и т. д.")
	fmt.Print("Введите целое число: ")
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Input error: ")
		return
	}
	if num%2 == 1 {
		if num > 0 {
			fmt.Println("Нечетное положительное число")
		} else {
			fmt.Println("Нечетное отрицательное число")
		}
	} else {
		if num > 0 {
			fmt.Println("Четное положительное число")
		} else if num < 0 {
			fmt.Println("Четное отрицательное число")
		} else {
			fmt.Println("Нулевое число")
		}
	}
}
