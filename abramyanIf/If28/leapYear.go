package If28

import "fmt"

func IsLeapYear() {
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
