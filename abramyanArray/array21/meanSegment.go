package array21

import "fmt"

func MeanSegment() float64 {
	fmt.Print("Задайте размер массива: ")
	var n, k, l int
	var mean, val, sum float64
	var arr []float64
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("input error: ", err)
		return 0
	}
	fmt.Println("Задайте элементы массива:")
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&val)
		if err != nil {
			fmt.Println("input error: ", err)
			return 0
		}
		arr = append(arr, val)
	}
	fmt.Print("Задайте отрезок среднее арифметическое которое мы вычислим(K < L): ")
	_, err = fmt.Scan(&k, &l)
	if err != nil {
		fmt.Println("input error: ", err)
		return 0
	}
	if k > l {
		fmt.Println("Cледуйте указаниям!")
		return 0
	}
	for i := k; i <= l; i++ {
		sum = sum + arr[i]
	}
	mean = sum / (float64(l - k + 1))
	return mean

}
