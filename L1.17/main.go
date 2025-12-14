package main

import (
	"cmp"
	"fmt"
)

func binarySearchGeneric[T cmp.Ordered](arr []T, target T) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		switch {
		case arr[mid] == target:
			return mid
		case arr[mid] < target:
			left = mid + 1
		default:
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	fmt.Println("Массив:", arr)

	testCases := []int{1, 7, 19, 10, 0, 20}
	fmt.Println("Ищем:", testCases)

	for _, target := range testCases {
		idx := binarySearchGeneric(arr, target)

		fmt.Printf("%d → %d\n",
			target, idx)
	}
}
