package main

func main() {
	// a := make([]int, 0)
	// fmt.Println("slice a is equal to nil:", a == nil)
	// fmt.Println(a)

	// var nums []int
	// fmt.Println(nums)
	// fmt.Printf("var nums is in address %p\n", &nums)
	// nums = append(nums, 1)
	// fmt.Println(nums)
	// fmt.Printf("var nums is in address %p\n", &nums)
	// nums = append(nums, 2)
	// fmt.Println(nums)
	// fmt.Printf("var nums is in address %p\n", &nums)

	// str := []string{"one", "two", "three"}
	// for i, v := range str {
	// 	fmt.Println("index:", i, "value:", v)
	// }

	// mySlice := []float64{1.2, 5.6}
	// // mySlice[2] = 6
	// mySlice = append(mySlice, 6)
	// a := 10
	// mySlice[0] = float64(a)
	// // mySlice[3] = 10.10
	// mySlice = append(mySlice, 10.10)
	// mySlice = append(mySlice, float64(a))
	// fmt.Println(mySlice)

	// nums := []float64{.1, .2, .3}
	// nums = append(nums, 10.1)
	// nums = append(nums, 4.1, 5.5, 6.6)
	// fmt.Println(nums)
	// n := []float64{1.1, 1.2}
	// nums = append(nums, n...)
	// fmt.Println(nums)

	// if len(os.Args)-1 <= 2 || len(os.Args)-1 >= 10 {
	// 	fmt.Println("Please give between 2 and 10 numbers")
	// } else {
	// 	fmt.Printf("os.Args is of type %T\n", os.Args)
	// 	nums := os.Args[1:]
	// 	sum, product := 0, 1
	// 	for _, str := range nums {
	// 		value, _ := strconv.Atoi(str)
	// 		sum += value
	// 		product *= value
	// 	}
	// 	fmt.Println("Sum:", sum, "Product:", product)
	// }

	// nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	// nums = nums[1 : len(nums)-2]
	// fmt.Println(nums)
	// // fmt.Println(len(nums), cap(nums))
	// // fmt.Println(nums[6:]) // Error
	// sum := 0
	// for _, v := range nums {
	// 	sum += v
	// }
	// fmt.Println("Sum:", sum)

	// friends := []string{"Marry", "John", "Paul", "Diana"}
	// newFriends := make([]string, len(friends))
	// _ = copy(newFriends, friends)
	// newFriends[0] = "Alice"
	// fmt.Println(friends, newFriends)

	// friends := []string{"Marry", "John", "Paul", "Diana"}
	// var newFriends []string
	// for i := range friends {
	// 	newFriends = append(newFriends, friends[i])
	// }
	// newFriends[0] = "Alice"
	// fmt.Println(friends, newFriends)

	// friends := []string{}
	// fmt.Println("var friends equals to nil:", friends == nil)
	// fmt.Printf("var friends is of length %d and capacity %d\n", len(friends), cap(friends))

	// years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	// newYears := append(years[:3], years[len(years)-3:]...)
	// fmt.Printf("%#v\n", newYears)
}
