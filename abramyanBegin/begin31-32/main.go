//Дано значение температуры T в градусах Фаренгейта. Определить
//значение этой же температуры в градусах Цельсия и наоборот.
//Температура по Цельсию TC и температура по Фаренгейту TF связаны следующим соотношением:
//TC = (TF − 32)·5/9

package main

import "fmt"

func main() {
	var temperature float64
	var measurement int8
	fmt.Println("В каких еденицах измеряется температура?\n1 - Цельсия;   2 - Фаренгейт")
	_, err := fmt.Scan(&measurement)
	if err != nil {
		fmt.Println("Input error: ", err)
		return
	}
	if measurement == 1 {
		fmt.Print("Температура в Цельсия = ")
		_, err := fmt.Scan(&temperature)
		if err != nil {
			fmt.Println("Input error: ", err)
			return
		}
		Tf := temperature*9/5 + 32
		fmt.Println("Температура в Фаренгейтах: ", Tf)
	} else if measurement == 2 {
		fmt.Print("Температура в Фаренгейт = ")
		_, err := fmt.Scan(&temperature)
		if err != nil {
			fmt.Println("Input error: ", err)
			return
		}
		Tc := (temperature - 32) * 5 / 9
		fmt.Println("Температура в Цельсиях: ", Tc)
	} else {
		fmt.Println("Пожалуйста используйте только 1 или 2")
	}
}
