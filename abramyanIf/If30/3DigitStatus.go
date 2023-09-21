package If30

import "fmt"

func ThreeDigitStatus() {
	fmt.Println(" Дано целое число, лежащее в диапазоне 1–999. Вывести его строкуописание вида «четное двузначное число», «нечетное трехзначное число»\nи т. д.")
	fmt.Print("Введите число от 1 до 999: ")
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Input error: ", err)
		return
	}
	if num < 1 || num > 999 {
		fmt.Println("Следуйте указаниям пожалуйста")
	} else {
		if num/100 != 0 {
			if num%2 == 1 {
				fmt.Println("Трехзначное нечетное число")
			} else {
				fmt.Println("Трехзначное четное число")
			}
		} else if num/10 != 0 {
			if num%2 == 1 {
				fmt.Println("Двузначное нечетное число")
			} else {
				fmt.Println("Двузначное четное число")
			}
		} else {
			if num%2 == 1 {
				fmt.Println("Однозначное нечетное число")
			} else {
				fmt.Println("Однозначное четное число")
			}
		}
	}
}
