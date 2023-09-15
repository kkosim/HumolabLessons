//Дано трехзначное число. Проверить истинность высказывания:
//«Все цифры данного числа различны»

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Введите трёхзначное число: ")
	//Unfortunately this error caption does not catch "wrong type" error
	//for example it reads normally these inputs: 514ijbg or 568.5642
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Input error: ", err)
		return
	} else if n > 999 || n < -999 {
		fmt.Println("Трёхзначное число!")
		return
	}
	if (n/100 != (n/10)%10) && ((n/10)%10 != n%10) && (n/100 != n%10) {
		fmt.Println("Все цифры данного числа различны")
	} else {
		fmt.Println("Цифры данного числа не различны")
	}
}
