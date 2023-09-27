package Leetcode

func KidsWithCandies(candies []int, extraCandies int) []bool {
	maxx := 0
	var ismax []bool
	for i := 0; i < len(candies); i++ {
		if candies[i] > maxx {
			maxx = candies[i]
		}
	}
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= maxx {
			ismax = append(ismax, true)
		} else {
			ismax = append(ismax, false)
		}
	}
	return ismax
}
