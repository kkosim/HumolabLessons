//Даны катеты прямоугольного треугольника a и b. Найти его гипотенузу c и периметр P

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Print("Введите значения катета и гипотенузы треугольника: ")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Ошибка при чтении из терминала", err)
		return
	}
	c := math.Sqrt(a*a + b*b)
	P := a + b + c
	fmt.Println("Гипотенуза треугольника: ", c)
	fmt.Println("Периметр треугольника: ", P)
}
