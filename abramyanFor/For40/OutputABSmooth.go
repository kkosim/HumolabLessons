package For40

import "fmt"

func ABSmooth() {
	fmt.Println("Даны целые числа A и B (A < B). Вывести все целые числа от A до B\nвключительно; при этом число A должно выводиться 1 раз, число A + 1\nдолжно выводиться 2 раза и т. д.")
	fmt.Print("Введите два числа A и В (А < B): ")
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Input error: ", err)
		return
	}
	if a >= b {
		fmt.Println("Следуйте указаниям")
	} else {
		for i := a; i <= b; i++ {
			for j := 1; j <= i-a+1; j++ {
				fmt.Print(i, " ")
			}
			fmt.Println()
		}
	}

}
