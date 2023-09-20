package floorAccess

import (
	"fmt"
	"strings"
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

	fmt.Print("What floor would you like to go? ")
	var floor int
	_, err := fmt.Scan(&floor)
	if err != nil {
		fmt.Println("Input error: ", err)
		return 0, fmt.Errorf("no such floor")
	}
	if floor >= 1 && floor <= 4 {
		return floor, nil
	} else {
		return 0, fmt.Errorf("no such floor")
	}
}

func CheckId(id string, floor int) {

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
	}
}
