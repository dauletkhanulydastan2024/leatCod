package main

import "fmt"

// func main() {
// 	nums := [...]int{1, 2, 3}

// 	fmt.Println(ans)
// }

// func getConcatenation(nums) []int {
// 	ans := make([]int, len(nums)*2)
//  copy(ans[:len(nums)], nums)
//  copy(ans[len(nums):], nums)
//  return ans
// }
func main() {
	// arr := [...]string{"A", "B", "S", "D", "F", "G"}
	// for i := 0; i < 4; i++ {
	// 	fmt.Println(arr[i])
	// }
	// arr := [...]int{1, 1, 1, 1, 1, 1, 1}
	// sum := 0
	// for i := 0; i < len(arr); i++ {
	// 	sum = sum + arr[i]
	// }
	// fmt.Println(sum)
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		continue
	// 	}

	// 	fmt.Println(i)
	// }
	// arr := [...]int{1, 1, 1, 1, 1, 1, 1}
	// for i := range arr {
	// 	fmt.Println(i)
	// }
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(i, arr[i])
	//
	nums := [3]int{1, 2, 3}
	nums2 := [6]int{}
	for i := 0; i < len(nums); i++ {
		nums2[i] = nums[i]
	}
	nums2[3] = nums[0]
	nums2[4] = nums[1]
	nums2[5] = nums[2]
	fmt.Println(nums2)
}
