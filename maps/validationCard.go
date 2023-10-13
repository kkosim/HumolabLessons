package maps

import (
	"fmt"
	"regexp"
)

var cardNumberRegex = `^\d{4}-\d{4}$`

type Card struct {
	CardNumber string
	CardHolder string
	Bank       string
}

func CardFill(num, holder string) Card {
	fmt.Print("Введите номер карты и Имя владельца: ")
	var currCard Card
	currCard.CardNumber = num
	currCard.CardHolder = holder
	_, err := fmt.Scan(&num, &holder)
	if err != nil {
		fmt.Println("input error: ", err)
		return Card{"", "", ""}
	}
	if num[0] >= 53 {
		currCard.Bank = "Humo bank"
	} else {
		currCard.Bank = "Other bank"
	}
	return currCard
}

func ValidationCard(List1 []string, List2 []string) {
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
	//see if you can update

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
