package Floor

import (
	"fmt"
)

type User struct {
	ID    string
	Floor int
}

func (u User) GetEmployeeID() (string, int) {

	fmt.Print("Input your ID: ")
	_, _ = fmt.Scan(&u.ID)
	switch u.ID[:1] {
	case "1":
		u.Floor = 1
	case "2":
		u.Floor = 2
	case "3":
		u.Floor = 3
	case "4":
		u.Floor = 4
	default:
		fmt.Println("Wrong ID!")
	}
	return u.ID, u.Floor
}

func GetWantedFloor() (int, error) {

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

func (u User) CheckId(floor int) {

	switch u.Floor {
	case 1:
		if floor > 1 {
			fmt.Println("Access denied!")
		} else {
			fmt.Println("Welcome!")
		}
	case 2:
		if floor > 2 {
			fmt.Println("Access denied!")
		} else {
			fmt.Println("Welcome!")
		}
	case 3:
		if floor > 3 {
			fmt.Println("Access denied!")
		} else {
			fmt.Println("Welcome!")
		}
	case 4:
		if floor > 4 {
			fmt.Println("Access denied!")
		} else {
			fmt.Println("Welcome Big Boss!")
		}
	default:
		fmt.Println("По идее unreachable code")
	}
}

func kosim() {
	var str string
	example(str)
}

func example(exp interface{}) {
	fmt.Println("SHIT: ", exp)
}
