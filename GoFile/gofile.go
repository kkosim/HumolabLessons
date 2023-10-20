package main

import (
	"fmt"
	"io"
	"os"
)

var Lorem = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin tellus ligula, feugiat ac mauris ac, tincidunt mollis leo. Donec mollis consequat lorem nec bibendum. Proin elementum, lacus non gravida"

func main() {
	//file, _ := os.Create("files/texxt.txt")
	//file.Close()
	LoremText := TwentyGap(Lorem)
	FillFile(LoremText)
	fmt.Println(FiletextToConsole())
	if LoremText == FiletextToConsole() {
		fmt.Println("Alike")
	} else {
		fmt.Println("Not Alike")
	}
}

func FillFile(text string) {
	file, err := os.Create("files/text.txt")
	if err != nil {
		fmt.Println("unable to open file: ", err.Error())
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error while closing file", err.Error())
		}
	}(file)
	_, err = file.WriteString(text);
	if err != nil {
		fmt.Println("unable to fill file: ", err.Error())
		os.Exit(2)
	}
}

func FiletextToConsole() (str string) {
	file, err := os.Open("files/text.txt")
	if err != nil {
		fmt.Println("error while opening file: ", err.Error())
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error while closing file: ", err.Error())
		}
	}(file)
	data := make([]byte, 250)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		str += string(data[:n])
	}
	return
}

func TwentyGap(LoremText string) string {
	str := ""
	bytes := []byte(LoremText)
	fmt.Println("jnr: ", len(bytes))
	for i := 0; i < len(bytes); i++ {
		if i%20 == 0 {
			str += "\n"
			fmt.Println("check: ", str)
		}
		str += string(bytes[i])
		fmt.Println("check2: ", str)
	}
	return str
}
