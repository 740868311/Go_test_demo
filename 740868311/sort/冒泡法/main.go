package main

import "fmt"

func main() {
	arr := []int{4, 1, 5, 7, 3, 2, 9}

	fmt.Println(maopao(arr))
}

func maopao(arr []int) []int {
	var swap int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				swap = arr[i]
				arr[i] = arr[j]
				arr[j] = swap
			}
		}
	}
	return arr
}
