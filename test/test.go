package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	start := time.Now()
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()
	go func() {
		for i := 5; i < 8; i++ {
			fmt.Println(i)
		}
	}()
	elapsedTime := time.Since(start)
	fmt.Println("Total Time For Execution: " + elapsedTime.String())
	time.Sleep(time.Second)
	fmt.Println("Finish? Enter y/n")
	input := ""
	for input != "y" && input != "n" {
		fmt.Scanln(&input)
		if input != "y" && input != "n" {
			fmt.Println("Неверно!")
			continue
		} else {
			break
		}
	}
	fmt.Println("The End")
}
