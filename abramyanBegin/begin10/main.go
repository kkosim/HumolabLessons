//Даны два ненулевых числа. Найти сумму, разность, произведение и
//частное их квадратов

package main

import "fmt"

func main() {
	var a, b float32
	fmt.Print("Введите два неотрицательных числа: ")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Ошибка при чтении данных: ", err)
		return
	}
	fmt.Println("Сумма квадратов чисел: ", (a*a)+(b*b))
	fmt.Println("Разность квадратов чисел: ", (a*a)-(b*b))
	fmt.Println("Произведение квадратов чисел: ", (a*a)*(b*b))
	fmt.Println("Частное квадратов чисел: ", (a*a)/(b*b))
}
