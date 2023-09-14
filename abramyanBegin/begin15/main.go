// Дана площадь S круга. Найти его диаметр D и длину L окружности,
//ограничивающей этот круг, учитывая, что L = 2·π·R, S = π·R^2.
//В качестве значения π использовать 3.14

package main

import (
	"fmt"
	"math"
)

func main() {
	var S float64
	const pi = 3.14
	fmt.Print("Площадь круга: ")
	_, err := fmt.Scan(&S)
	if err != nil {
		fmt.Println("Ошибка при чтении из терминала", err)
		return
	}
	R := math.Sqrt(S / pi)
	L := 2 * pi * R
	fmt.Printf("Диаметр круга: %v\nДлина круга: %v", 2*R, L)
}
