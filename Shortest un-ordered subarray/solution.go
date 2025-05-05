package main

import "fmt"

func increasing(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func decreasing(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			return false
		}
	}
	return true
}

func shortestUnorderedSubarray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if increasing(arr) || decreasing(arr) {
		return 0
	}
	return 3
}

func main() {
	fmt.Println(shortestUnorderedSubarray([]int{7, 9, 10, 8, 11}))
	fmt.Println(shortestUnorderedSubarray([]int{5, 4, 3, 2, 1}))
}
