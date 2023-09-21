package For39

import "fmt"

func ABTimes() {
	fmt.Println("Даны целые положительные числа A и B (A < B). Вывести все целые\nчисла от A до B включительно; при этом каждое число должно выводиться\nстолько раз, каково его значение (например, число 3 выводится 3 раза).")
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
			for j := 1; j <= i; j++ {
				fmt.Print(i, " ")
			}
			fmt.Println()
		}
	}

}
