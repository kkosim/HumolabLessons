//Программа берет в качестве данных координаты шахматной фигуры
//и проверяет может ли данная фигура "перепрыгнуть" на новую координату

package main

import "fmt"

func main() {
	var x1, y1, x2, y2, chessPiece uint8
	fmt.Println("Какая фигура сейчас на шахматной доске?")
	fmt.Println("1 - Пешка\n2 - Ладья\n3 - Конь\n4 - Слон\n5 - Король\n6 - Ферзь")
	_, err := fmt.Scan(&chessPiece)
	if err != nil {
		fmt.Print("Input error", err)
		return
	}
	switch chessPiece {

	case 1:
		fmt.Print("Введите текущие координаты пешки: ")
		_, err := fmt.Scan(&x1, &y1)
		if err != nil {
			fmt.Print("Input error", err)
			return
		}
		if (x1 == 0 || x1 > 8) || (y1 == 0 || y1 > 8) {
			fmt.Print("Таких координат нет!")
			return
		}
		fmt.Print("Введите координаты будущей позиции пешки: ")
		_, err = fmt.Scan(&x2, &y2)
		if err != nil {
			fmt.Print("Input error", err)
			return
		}
		if (x2 == 0 || x2 > 8) || (y2 == 0 || y2 > 8) {
			fmt.Print("Таких координат нет!")
			return
		}
		if y2-y1 == 1 {
			fmt.Print("Пешка может за один ход перейти на данную позицию")
		} else if x1 == x2 {
			fmt.Print("Пешка не может перейти на данную позицию")
		} else {
			fmt.Print("Пешка не может за один ход перейти на данную позицию")
		}

	case 2:
		fmt.Print("Введите текущие координаты ладьи: ")
		_, err := fmt.Scan(&x1, &y1)
		if err != nil {
			fmt.Print("Input error", err)
			return
		}
		if (x1 == 0 || x1 > 8) || (y1 == 0 || y1 > 8) {
			fmt.Print("Таких координат нет!")
			return
		}
		fmt.Print("Введите координаты будущей позиции ладьи: ")
		_, err = fmt.Scan(&x2, &y2)
		if err != nil {
			fmt.Print("Input error", err)
			return
		}
		if (x2 == 0 || x2 > 8) || (y2 == 0 || y2 > 8) {
			fmt.Print("Таких координат нет!")
			return
		}
		if x1 == x2 || y1 == y2 {
			fmt.Print("Ладья может за один ход перейти на данную позицию")
		} else {
			fmt.Print("Ладья не может за один ход перейти на данную позицию")
		}

	case 3:
		fmt.Print("Введите текущие координаты коня: ")
		_, err := fmt.Scan(&x1, &y1)
		if err != nil {
			fmt.Print("Input error", err)
			return
		}
		if (x1 == 0 || x1 > 8) || (y1 == 0 || y1 > 8) {
			fmt.Print("Таких координат нет!")
			return
		}
		fmt.Print("Введите координаты будущей позиции коня: ")
		_, err = fmt.Scan(&x2, &y2)
		if err != nil {
			fmt.Print("Input error", err)
			return
		}
		if (x2 == 0 || x2 > 8) || (y2 == 0 || y2 > 8) {
			fmt.Print("Таких координат нет!")
			return
		}
		if (x1+2 == x2 && y1+1 == y2) || (x1+1 == x2 && y1+2 == y2) || (x1+2 == x2 && y1-1 == y2) || (x1+1 == x2 && y1-2 == y2) || (x1-2 == x2 && y1+1 == y2) || (x1-1 == x2 && y1+2 == y2) || (x1-2 == x2 && y1-1 == y2) || (x1-1 == x2 && y1-2 == y2) {
			fmt.Print("Конь может за один ход перейти на данную позицию")
		} else {
			fmt.Print("Конь не может за один ход перейти на данную позицию")
		}

	case 4:
		fmt.Print("Введите текущие координаты слона: ")
		_, err := fmt.Scan(&x1, &y1)
		if err != nil {
			fmt.Print("Input error", err)
			return
		}
		if (x1 == 0 || x1 > 8) || (y1 == 0 || y1 > 8) {
			fmt.Print("Таких координат нет!")
			return
		}
		fmt.Print("Введите координаты будущей позиции слона: ")
		_, err = fmt.Scan(&x2, &y2)
		if err != nil {
			fmt.Print("Input error", err)
			return
		}
		if (x2 == 0 || x2 > 8) || (y2 == 0 || y2 > 8) {
			fmt.Print("Таких координат нет!")
			return
		}
		if x2-x1 == y2-y1 || x1-x2 == y2-y1 {
			fmt.Print("Слон может за один ход перейти на данную позицию")
		} else if (x2-x1)%2 != (y2-y1)%2 || (x2-x1)%2 != (y1-y2)%2 {
			fmt.Print("Слон не может перейти на данную позицию")
		} else {
			fmt.Print("Слон не может за один ход перейти на данную позицию")
		}

	case 5:
		fmt.Print("Введите текущие координаты короля: ")
		_, err := fmt.Scan(&x1, &y1)
		if err != nil {
			fmt.Print("Input error", err)
			return
		}
		if (x1 == 0 || x1 > 8) || (y1 == 0 || y1 > 8) {
			fmt.Print("Таких координат нет!")
			return
		}
		fmt.Print("Введите координаты будущей позиции короля: ")
		_, err = fmt.Scan(&x2, &y2)
		if err != nil {
			fmt.Print("Input error", err)
			return
		}
		if (x2 == 0 || x2 > 8) || (y2 == 0 || y2 > 8) {
			fmt.Print("Таких координат нет!")
			return
		}
		if (x1 == x2 && y1+1 == y2) || (x1+1 == x2 && y1+1 == y2) || (x1+1 == x2 && y1 == y2) || (x1+1 == x2 && y1-1 == y2) || (x1 == x2 && y1-1 == y2) || (x1-1 == x2 && y1-1 == y2) || (x1-1 == x2 && y1 == y2) || (x1-1 == x2 && y1+1 == y2) {
			fmt.Print("Король может за один ход перейти на данную позицию")
		} else {
			fmt.Print("Король не может за один ход перейти на данную позицию")
		}

	case 6:
		fmt.Print("Введите текущие координаты ферзя: ")
		_, err := fmt.Scan(&x1, &y1)
		if err != nil {
			fmt.Print("Input error", err)
			return
		}
		if (x1 == 0 || x1 > 8) || (y1 == 0 || y1 > 8) {
			fmt.Print("Таких координат нет!")
			return
		}
		fmt.Print("Введите координаты будущей позиции ферзя: ")
		_, err = fmt.Scan(&x2, &y2)
		if err != nil {
			fmt.Print("Input error", err)
			return
		}
		if (x2 == 0 || x2 > 8) || (y2 == 0 || y2 > 8) {
			fmt.Print("Таких координат нет!")
			return
		}
		if (x2-x1 == y2-y1 || x1-x2 == y2-y1) || (x1 == x2 || y1 == y2) {
			fmt.Print("Ферзь может за один ход перейти на данную позицию")
		} else {
			fmt.Print("Ферзь не может за один ход перейти на данную позицию")
		}
	default:
		fmt.Println("Ты чё курил?")
		return
	}
}
