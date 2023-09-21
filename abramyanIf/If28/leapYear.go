package If28

import "fmt"

func IsLeapYear() {
	fmt.Println("Дан номер года (положительное целое число). Определить количество\nдней в этом году, учитывая, что обычный год насчитывает 365 дней, а\nвисокосный — 366 дней. Високосным считается год, делящийся на 4, за\nисключением тех годов, которые делятся на 100 и не делятся на 400\n(например, годы 300, 1300 и 1900 не являются високосными, а 1200 и 2000\n— являются)")
	fmt.Print("Введите год: ")
	var year int
	_, err := fmt.Scan(&year)
	if err != nil {
		fmt.Println("Input error: ")
		return
	}
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				fmt.Println("Дней в году 366")
				return
			}
			fmt.Println("Дней в году 365")
			return
		}
		fmt.Println("Дней в году 366")
		return
	} else {
		fmt.Println("Дней в году 365")
		return
	}
}
