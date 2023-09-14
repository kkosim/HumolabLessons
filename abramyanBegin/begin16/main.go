//Найти расстояние между двумя точками с заданными координатами x1 и x2 на числовой оси: |x2 − x1|.

package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2 float64
	fmt.Print("Координаты точек х1 и х2 на числовой оси: ")
	_, err := fmt.Scan(&x1, &x2)
	if err != nil {
		fmt.Println("Ошибка при чтении из терминала: ", err)
		return
	}
	fmt.Println("Растоние между точками: ", math.Abs(x1-x2))
}
