package main

import "fmt"

func main() {

	arr := []int{4, 1, 5, 7, 3, 2, 9}

	fmt.Println(freeSort(arr))
}

func freeSort(arr []int) []int {
	sign := arr[0]

	var leftA []int
	var rightA []int
	for _, v := range arr {
		if v > sign {
			rightA = append(rightA, v)
		} else if v < sign {
			leftA = append(leftA, v)
		}
	}

	if len(leftA) > 0 {
		leftA = freeSort(leftA)
	}

	if len(rightA) > 0 {
		rightA = freeSort(rightA)
	}

	var res []int
	res = append(leftA, sign)
	res = append(res, rightA...)
	return res
}
