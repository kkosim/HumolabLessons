//Даны два круга с общим центром и радиусами R1 и R2 (R1 > R2).
//Найти площади этих кругов S1 и S2, а также площадь S3 кольца, внешний
//радиус которого равен R1, а внутренний радиус равен R2
//В качестве значения π использовать 3.14

package main

import "fmt"

func main() {
	var r1, r2 float64
	const pi = 3.14
	fmt.Print("Введите значение радиуса 1го круга и 2го круга: ")
	_, err := fmt.Scan(&r1, &r2)
	if err != nil {
		fmt.Println("Ошибка при чтении из терминала: ", err)
		return
	} else if r1 < r2 {
		fmt.Println("Второй круг вписан в первый по условию задачи!")
		return
	}
	s1 := pi * r1 * r1
	s2 := pi * r2 * r2
	s3 := s1 - s2
	fmt.Printf("Площадь 1го круга: %v\nПлощадь 2го круга: %v\nПлощадь кольца: %v", s1, s2, s3)
}
