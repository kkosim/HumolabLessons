package floorAccess

import (
	"fmt"
)

func GetEmployeeID() (id string, err error) {

	fmt.Print("Input your ID: ")
	_, _ = fmt.Scan(&id)
	if id[0:1] == "1" || id[0:1] == "2" || id[0:1] == "3" || id[0:1] == "4" {
		return id, nil
	} else {
		return "", fmt.Errorf("wrong id")
	}
}

func GetFloor() (int, error) {

	fmt.Print("Which floor do you need to go to?: ")
	var floor int
	_, err := fmt.Scan(&floor)
	if err != nil {
		fmt.Println("Input error: ", err)
		return 0, fmt.Errorf("wrong type")
	}
	if floor >= 1 && floor <= 4 {
		return floor, nil
	} else {
		return 0, fmt.Errorf("no such floor")
	}
}

func CheckId(id string, floor int) {

	switch id[0:1] {
	case "1":
		if floor > 1 {
			fmt.Println("Access denied!")
		} else {
			fmt.Println("Welcome!")
		}
	case "2":
		if floor > 2 {
			fmt.Println("Access denied!")
		} else {
			fmt.Println("Welcome!")
		}
	case "3":
		if floor > 3 {
			fmt.Println("Access denied!")
		} else {
			fmt.Println("Welcome!")
		}
	case "4":
		if floor > 4 {
			fmt.Println("Access denied!")
		} else {
			fmt.Println("Welcome Big Boss!")
		}
	default:
		fmt.Println("По идее unreachable code")
	}
}
