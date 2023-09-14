//Даны три точки A, B, C на числовой оси. Найти длины отрезков AC
//и BC и их сумму

package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C float64
	fmt.Print("Координаты точек A, B, C на числовой оси: ")
	_, err := fmt.Scan(&A, &B, &C)
	if err != nil {
		fmt.Println("Ошибка при чтении из терминала: ", err)
		return
	}
	AC := math.Abs(A - C)
	BC := math.Abs(B - C)
	fmt.Printf("Длина отрезка АС: %v, отрезка ВС: %v, сумма этих отрезков: %v", AC, BC, AC+BC)
}
