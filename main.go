package main

import (
	"fmt"
)

func main() {
	//var somoni float32
	var somoni int
	fmt.Print("How much somoni you have: ")
	_, err := fmt.Scan(&somoni)
	if err != nil {
		fmt.Println("Ошибка при чтении из терминала:", err.Error())
		return
	}

	//Converter in golang
	/*integer, err := strconv.Atoi("8")
	if err != nil {
		fmt.Println("Ошибка конвертации:", err.Error())
		return
	}
	fmt.Println(integer)

	*/
	//Contains in golang
	/*
		if strings.Contains("Вот какое-то число", "число") {
			fmt.Println("contains")
		} else {
			fmt.Println("doesn't")
		}
	*/
	//newString := strings.Replace()
	//var rubles = somoni * 8.95
	//var dollars = somoni * 0.091
	var rubles = somoni * 895 / 100
	var dollars = somoni * 91 / 1000
	fmt.Printf("Your money is as valuable as: %v rubles\n", rubles)
	fmt.Printf("Your money is as valuable as: %v dollars\n", dollars)
}
