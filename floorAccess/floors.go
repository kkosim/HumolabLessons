package floorAccess

import (
	"fmt"
	"strings"
)

func GetEmployeeID() {
	fmt.Print("Введите ID: ")
	var id string
	_, err := fmt.Scan(&id)
	if err != nil {
		fmt.Println("Input error: ", err.Error())
		return
	}
	fmt.Print("На какой этаж вы хотите попасть? ")
	var floor int
	_, err = fmt.Scan(&floor)
	if err != nil {
		fmt.Println("Input error: ", err)
		return
	}
	if floor < 1 || floor > 4 {
		fmt.Print("Шутник")
		return
	}
	if strings.HasPrefix(id, "1") {
		if floor > 1 {
			fmt.Println("Access denied!")
		} else {
			fmt.Println("Welcome!")
		}
	} else if strings.HasPrefix(id, "2") {
		if floor > 2 {
			fmt.Println("Access denied!")
		} else {
			fmt.Println("Welcome!")
		}
	} else if strings.HasPrefix(id, "3") {
		if floor > 3 {
			fmt.Println("Access denied!")
		} else {
			fmt.Println("Welcome!")
		}
	} else if strings.HasPrefix(id, "4") {
		if floor > 4 {
			fmt.Println("Access denied!")
		} else {
			fmt.Println("Welcome Big Boss!")
		}
	} else {
		fmt.Println("Wrong ID!")
	}
	return
}
