package main

import (
	"fmt"
	"math/rand/v2"
)

// используем середину массива в качестве опорного элемента и рекурсию
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]

	var left, middle, right []int

	for _, num := range arr {
		switch {
		case num < pivot:
			left = append(left, num)
		case num > pivot:
			right = append(right, num)
		default:
			middle = append(middle, num)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	result := append(left, middle...)
	result = append(result, right...)

	return result
}

// генерация случайного массива
func generateRandomArray(size, maxValue int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.IntN(maxValue)
	}
	return arr
}

func main() {
	arr := generateRandomArray(15, 100)
	fmt.Println("До сортировки:", arr)
	sorted := quickSort(arr)
	fmt.Println("После сортировки:", sorted)
}
