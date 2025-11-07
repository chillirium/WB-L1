package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Print("Первое число: ")
	fmt.Scan(&a)
	fmt.Print("Второе число: ")
	fmt.Scan(&b)

	fmt.Println("Поначалу:", a, b)

	a ^= b
	b ^= a
	a ^= b

	fmt.Println("Опосля:", a, b)
}
