package maps

import (
	"fmt"
	"regexp"
)

var cardNumberRegex = `^\d{4}-^\d{4}$`

func ValidationCard() {
	var List1 = []string{"1222-3243", "6123-2321", "1112-2341", "7112-2344", "3211-2231"}
	var List2 = []string{"1222-3243", "1123-2321", "9112-2341", "2112-2344", "5211-2231"}
	mapListHumo := make(map[string]string)
	mapListOther := make(map[string]string)
	var List12 = append(List1, List2...)
	mapAllCards := make(map[string]string)
	for _, v := range List12 {
		mapAllCards[v] = ""
		if v[0] >= 53 {
			mapListHumo[v] = ""
		} else {
			mapListOther[v] = ""
		}
	}
	fmt.Print("Input card number: ")
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("input error: ", err)
		return
	}
	match, err := regexp.MatchString(cardNumberRegex, input)
	if err != nil {
		fmt.Println("regex error: ")
		return
	}
	if !match {
		fmt.Println("Such card number does not exist!")
		return
	}
	_, exist := mapAllCards[input]
	if exist {
		fmt.Println("Card already exists!")
	} else {
		fmt.Println("New card added!")
		if input[0] >= 53 {
			mapListHumo[input] = ""
		} else {
			mapListOther[input] = ""
		}
	}
	fmt.Printf("Humo cards: %v\nOther bank cards: %v", mapListHumo, mapListOther)
}
