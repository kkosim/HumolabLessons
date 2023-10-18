package main

import (
	"fmt"
	"sync"
	"time"
)

type Card struct {
	ID          int
	CardNumber  string
	ExpiryMonth string
	ExpiryYear  string
	CVV         string
	Balance     int
	IsBlocked   bool
}

type Product struct {
	ID    int
	Name  string
	Price int
}

var cards = map[int]Card{
	1: {
		ID:          1,
		CardNumber:  "1234567890123456",
		ExpiryMonth: "01",
		ExpiryYear:  "23",
		CVV:         "123",
		Balance:     1000,
		IsBlocked:   false,
	},
	2: {
		ID:          2,
		CardNumber:  "9876543210987654",
		ExpiryMonth: "05",
		ExpiryYear:  "24",
		CVV:         "321",
		Balance:     500,
		IsBlocked:   false,
	},
}

var products = map[int]Product{
	1: {
		ID:    1,
		Name:  "HeadPhone",
		Price: 80,
	},
	2: {
		ID:    2,
		Name:  "PowerBank",
		Price: 200,
	},
	3: {
		ID:    3,
		Name:  "Smartphone",
		Price: 700,
	},
}

var mutex = sync.Mutex{}

func main() {
	go scheduler()
	go buyProduct()
	showProducts()

	time.Sleep(10 * time.Second)
}

func scheduler() {
	for {
		time.Sleep(2 * time.Second)
		showProducts()
	}
}

func showProducts() {
	fmt.Println("Products:")
	for _, product := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %d\n", product.ID, product.Name, product.Price)
	}
	fmt.Println()
}

func buyProduct() {
	for {
		var productID int
		var cardID int
		var confirm string

		fmt.Print("Enter the ID of the product you want to buy: ")
		_, err := fmt.Scanln(&productID)
		if err != nil {
			return
		}

		fmt.Print("Enter the ID of the card for payment: ")
		_, err = fmt.Scanln(&cardID)
		if err != nil {
			return
		}

		product, ok := products[productID]
		if !ok {
			fmt.Println("Invalid product ID")
			continue
		}

		card, ok := cards[cardID]
		if !ok {
			fmt.Println("Invalid card ID")
			continue
		}

		if card.IsBlocked {
			fmt.Println("Card is blocked")
			continue
		}

		if card.Balance < product.Price {
			fmt.Println("Insufficient balance")
			continue
		}

		fmt.Printf("Confirm purchase of product '%s'? (yes/no): ", product.Name)
		_, err = fmt.Scanln(&confirm)
		if err != nil {
			return
		}

		if confirm == "yes" {
			mutex.Lock()
			card.Balance -= product.Price
			mutex.Unlock()

			fmt.Println("Purchase successful")
		} else {
			fmt.Println("Purchase canceled")
		}

		fmt.Println()
	}
}

//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//type Card struct {
//	number  string
//	balance float64
//	mu      sync.Mutex
//}
//
//type Product struct {
//	name  string
//	price float64
//}
//
//var cards = []Card{
//	{number: "1111 1111 1111 1111", balance: 1000},
//	{number: "2222 2222 2222 2222", balance: 2000},
//	{number: "3333 3333 3333 3333", balance: 3000},
//}
//
//var products = []Product{
//	{name: "Vaccum Cleaner", price: 200},
//	{name: "Smatrphone", price: 500},
//	{name: "Laptop", price: 900},
//}
//
//func main() {
//	go displayProducts()
//
//	time.Sleep(200 * time.Millisecond)
//
//	buyProduct(0, 0)
//
//	time.Sleep(10 * time.Second)
//}
//
//func displayProducts() {
//	ticker := time.NewTicker(time.Second)
//
//	for range ticker.C {
//		fmt.Println("========== Products List ==========")
//		for i, product := range products {
//			fmt.Printf("%d. %s - $%.2f\n", i+1, product.name, product.price)
//		}
//		fmt.Println("===================================")
//	}
//}
//
//func buyProduct(cardIndex, productIndex int) {
//	card := &cards[cardIndex]
//	product := &products[productIndex]
//
//	card.mu.Lock()
//	defer card.mu.Unlock()
//
//	if card.balance >= product.price {
//		card.balance -= product.price
//		fmt.Printf("Product \"%s\" purchased successfully with Card \"%s\"\n", product.name, card.number)
//		fmt.Printf("Remaining balance on Card \"%s\": $%.2f\n", card.number, card.balance)
//	} else {
//		fmt.Printf("Insufficient balance on Card \"%s\" to purchase Product \"%s\"\n", card.number, product.name)
//	}
//}
//