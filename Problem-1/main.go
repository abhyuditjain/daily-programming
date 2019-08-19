/**
Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

Bonus: Can you do this in one pass?
*/
package main

import "fmt"

func main() {
	arr := []int{10, 15, 3, 7}
	sum := 17
	fmt.Println(pairSum(arr, sum))
	sum = 12
	fmt.Println(pairSum(arr, sum))
}

func pairSum(arr []int, sum int) bool {
	complements := make(map[int]bool)

	for _, v := range arr {
		if complements[v] == true {
			return true
		}

		complements[sum-v] = true
	}
	return false
}
