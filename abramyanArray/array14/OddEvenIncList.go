package array14

import "fmt"

func OddEvenIncList() []int {
	fmt.Print("Сколько чисел вы хотите внести?  ")
	var n, val int
	var arr []int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("input error: ", err)
		return nil
	}
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&val)
		arr = append(arr, val)
		if err != nil {
			fmt.Println("input error: ", err)
			return nil
		}
	}
	var arr1, arr2 []int
	for i := 0; i < len(arr); i += 2 {
		arr1 = append(arr1, arr[i])
	}
	for i := 1; i < len(arr); i += 2 {
		arr2 = append(arr2, arr[i])
	}
	arr1 = append(arr1, arr2...)
	return arr1
}
