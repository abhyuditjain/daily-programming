/**
Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].

Follow-up: what if you can't use division?
*/
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(getArrayOfProductsUsingDivision(arr))
	fmt.Println(getArrayOfProductsWithoutUsingDivision(arr))
	arr = []int{3, 2, 1}
	fmt.Println(getArrayOfProductsUsingDivision(arr))
	fmt.Println(getArrayOfProductsWithoutUsingDivision(arr))
	arr = []int{1, 2, 0, 5}
	fmt.Println(getArrayOfProductsUsingDivision(arr))
	fmt.Println(getArrayOfProductsWithoutUsingDivision(arr))
	arr = []int{1, 2, 0, 0}
	fmt.Println(getArrayOfProductsUsingDivision(arr))
	fmt.Println(getArrayOfProductsWithoutUsingDivision(arr))
}

func getArrayOfProductsUsingDivision(arr []int) []int {
	nonZeroProduct := 1
	numZeros := 0

	for _, v := range arr {
		if v != 0 {
			nonZeroProduct *= v
		}
		if v == 0 {
			numZeros++
		}
	}

	newArr := make([]int, len(arr))

	if numZeros > 1 {
		return newArr
	}

	for i, v := range arr {
		if numZeros == 0 {
			newArr[i] = nonZeroProduct / v
		} else if v == 0 {
			newArr[i] = nonZeroProduct
		}
	}

	return newArr
}

func getArrayOfProductsWithoutUsingDivision(arr []int) []int {
	leftProduct := make([]int, len(arr))
	rightProduct := make([]int, len(arr))

	leftProduct[0] = 1
	rightProduct[len(rightProduct)-1] = 1

	for i := 1; i < len(arr); i++ {
		leftProduct[i] = leftProduct[i-1] * arr[i-1]
	}
	for i := len(arr) - 2; i >= 0; i-- {
		rightProduct[i] = rightProduct[i+1] * arr[i+1]
	}

	newArr := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		newArr[i] = leftProduct[i] * rightProduct[i]
	}

	return newArr
}
