package main

import "HumolabLessons/interfacePlanets"

func main() {
	/*var titanW GravityFinder = Titan{Name: "Titan"}
	  var proximaCentauriB GravityFinder = ProximaCentauriB{Name: "PCB"}
	  titanW.calculateWeight()
	  proximaCentauriB.calculateWeight()*/
	var (
		planet interfacePlanets.ProximaCentauriB
		moon   interfacePlanets.Titan
	)
	planet.Name = "PCB"
	moon.Name = "Titan"
	interfacePlanets.Weight(planet)
	interfacePlanets.Weight(moon)
	/*	var u floorAccess2.User
		u.ID, u.Floor = floorAccess2.User{}.GetEmployeeID()
		//id, err := floorAccess.GetEmployeeID()

		floor, _ := floorAccess2.GetWantedFloor()

		floorAccess2.User{u.ID = "1234", u.Floor = 1}.CheckId(floor)

	}*/

	/*var k float64
	  for i := 0; i < 10000; i++ {
	  	intSlice = append(intSlice, rand.Int())
	  	k += float64(intSlice[i])
	  }

	  sportsMan.GetCharacteristics()

	  fmt.Println(k)
	  avg := k / 10000
	  var avgindex int
	  timeStart := time.Now()
	  sort.Slice(intSlice, func(first, last int) bool { return intSlice[first] < intSlice[last] })
	  for i, v := range intSlice {
	  	if avg >= float64(v) {
	  		avgindex = i
	  	}
	  }
	  fmt.Printf("Average value: %v, %v numbers are greater than average\n", avg, avgindex)
	  fmt.Println(time.Since(timeStart))
	  timeStart2 := time.Now()
	  sort.Slice(intSlice, func(first, last int) bool { return intSlice[first] > intSlice[last] })
	  fmt.Println(time.Since(timeStart2))
	  sort.Slice(stringSlice, func(first, last int) bool { return stringSlice[first] < stringSlice[last] })
	  fmt.Println(stringSlice)

	  m := map[int]int{0: 0, 1: 10}

	  fmt.Println(m, m[0], m[1], m[2])*/

	/*var List1 = []string{"1222-3243", "6123-2321", "1112-2341", "7112-2344", "3211-2231"}
	  var List2 = []string{"1222-3243", "1123-2321", "9112-2341", "2112-2344", "5211-2231"}

	  maps.ValidationCard(List1, List2)
	*/
	//rand.Seed()

	//fmt.Println(Leetcode.ReverseVowels("leetcode"))

	/*candies := []int{2, 3, 5, 1, 3}
	  extraCandies := 3
	  fmt.Println(Leetcode.KidsWithCandies(candies, extraCandies))
	*/

	/*fish := "Разнообразный и богатый опыт говорит нам, что новая модель организационной деятельности позволяет выполнить важные задания по разработке прогресса профессионального сообщества. В частности, высококачественный прототип будущего проекта предполагает независимые способы реализации позиций, занимаемых участниками в отношении поставленных задач. Однозначно, ключевые особенности структуры проекта будут разоблачены. Мы вынуждены отталкиваться от того, что высокотехнологичная концепция общественного уклада играет важную роль в формировании первоочередных требований."
	  rune2.TwentyGap(fish)
	*/

	/*var nums []int
	  var n int
	  fmt.Print("Сколько нечетных чисел от 1 печатать? ")
	  _, err := fmt.Scan(&n)
	  if err != nil {
	  	fmt.Println("input error: ", err)
	  	return
	  }
	  for i := 1; i <= n*2-1; i += 2 {
	  	nums = append(nums, i)
	  }
	  fmt.Println(nums)

	*/
	/*fmt.Println("Вам какую задачу?")
	  fmt.Println("1 - If28\n2 - If29\n3 - If30\n4 - For39\n5 - For40")
	  var problem uint8
	  _, err := fmt.Scan(&problem)
	  if err != nil {
	  	fmt.Println("Input error: ", err)
	  	return
	  }
	  switch problem {
	  case 1:
	  	If28.IsLeapYear()
	  case 2:
	  	If29.EvenOddPosNeg()
	  case 3:
	  	If30.ThreeDigitStatus()
	  case 4:
	  	For39.ABTimes()
	  case 5:
	  	For40.ABSmooth()
	  default:
	  	fmt.Println("Следуйте указаниям!")
	  }*/

	//var somoni float32
	/*var somoni int
	  fmt.Print("How much somoni you have: ")
	  _, err := fmt.Scan(&somoni)
	  if err != nil {
	  	fmt.Println("Ошибка при чтении из терминала:", err.Error())
	  	return
	  }
	*/
	//Converter in golang
	/*integer, err := strconv.Atoi("8")
	  if err != nil {
	  	fmt.Println("Ошибка конвертации:", err.Error())
	  	return
	  }
	  fmt.Println(integer)

	*/
	//Contains in golang
	/*
		if strings.Contains("Вот какое-то число", "число") {
			fmt.Println("contains")
		} else {
			fmt.Println("doesn't")
		}
	*/
	//newString := strings.Replace()
	//var rubles = somoni * 8.95
	//var dollars = somoni * 0.091
	/*var rubles = somoni * 895 / 100
	  var dollars = somoni * 91 / 1000
	  fmt.Printf("Your money is as valuable as: %v rubles\n", rubles)
	  fmt.Printf("Your money is as valuable as: %v dollars\n", dollars)
	*/

}
