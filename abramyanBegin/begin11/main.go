//Даны два ненулевых числа. Найти сумму, разность, произведение и
//частное их модулей

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Print("Введите два неотрицательных числа: ")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Ошибка при чтении данных: ", err)
		return
	}
	fmt.Println("Сумма модулей чисел: ", math.Abs(a)+math.Abs(b))
	fmt.Println("Разность модулей чисел: ", math.Abs(a)-math.Abs(b))
	fmt.Println("Произведение модулей чисел: ", math.Abs(a)*math.Abs(b))
	fmt.Println("Частное модулей чисел: ", math.Abs(a)/math.Abs(b))
}
