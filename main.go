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
	arr := [...]int{1, 1, 1, 1, 1, 1, 1}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	fmt.Println(sum)
}
