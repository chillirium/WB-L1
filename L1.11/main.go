// go run main.go
// подразумеваются результат без повторений
package main

import "fmt"

func intersection[T comparable](a, b []T) []T {
	setA := make(map[T]bool)
	for _, v := range a {
		setA[v] = true
	}

	var result []T
	seen := make(map[T]bool)
	for _, v := range b {
		if setA[v] && !seen[v] {
			result = append(result, v)
			seen[v] = true
		}
	}
	return result
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}
	fmt.Println("Пересечение = ", intersection(A, B))
	C := []string{"apple", "banana", "cherry"}
	D := []string{"banana", "date", "apple", "elderberry"}
	fmt.Println("Пересечение для строк = ", intersection(C, D))
	E := []int{}
	F := []int{1, 2, 3, 2, 1}
	fmt.Println("Пересечение с пустым =", intersection(E, F))
	fmt.Println("Пересечение с собой =", intersection(F, F))
}
