package main

import "fmt"

func largest(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num >= max {
			max = num
		}
	}
	return max
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 111, 1111, 11111, 12, 111111}
	fmt.Println(largest(arr))
}
