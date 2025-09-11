// go run main.go
package main

import (
	"fmt"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	results := make(chan int, len(numbers)) //создание канала

	for _, num := range numbers {
		go func(n int) { //запуск горутин для каждого числа
			results <- n * n //запись в канал квадрата числа
		}(num)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(<-results) //вывод из канала
	}

	close(results) //закрытие канала
}
