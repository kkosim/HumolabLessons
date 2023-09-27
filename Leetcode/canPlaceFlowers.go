package Leetcode

func CanPlaceFlowers(flowerbed []int, n int) bool {
	k := 0
	length := len(flowerbed)
	if n > 0 {
		if length == 1 {
			if flowerbed[0] == 1 {
				return false
			} else {
				return true
			}
		}
		if length == 2 {
			if flowerbed[0] == 1 || flowerbed[1] == 1 || n == 2 {
				return false
			} else {
				return true
			}
		} else {

			for i := 0; i < length; i++ {
				if i == 0 {
					if flowerbed[0] == 0 && flowerbed[1] == 0 {
						k++
						flowerbed[i] = 1
					}
				} else if i == length-1 {
					if flowerbed[length-1] == 0 && flowerbed[length-2] == 0 {
						k++
						flowerbed[i] = 1
					}
				} else {
					if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
						k++
						flowerbed[i] = 1
					}
				}
			}
			if n > k {
				return false
			} else {
				return true
			}
		}
	} else {
		return true
	}
}
