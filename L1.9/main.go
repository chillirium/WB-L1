// go run main.go
package main

import (
	"fmt"
)

func generateNumbers(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, num := range numbers {
			out <- num
		}
	}()
	return out
}

func processNumbers(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num * 2
		}
	}()
	return out
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	inputChan := generateNumbers(numbers)
	outputChan := processNumbers(inputChan)

	for result := range outputChan {
		fmt.Println(result)
	}
}
