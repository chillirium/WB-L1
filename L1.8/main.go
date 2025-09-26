/*
go run main.go <число в 10- системе> <индекс бита в 1-инднксации справа налево> <значение>
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Использование: go run main.go <число> <индекс_бита> <значение>")
		fmt.Println("Пример: go run main.go 5 1 0")
		os.Exit(1)
	}

	num, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Printf("Ошибка парсинга числа: %v\n", err)
		os.Exit(1)
	}

	bitIndex, err := strconv.ParseUint(os.Args[2], 10, 64)
	if err != nil || bitIndex < 1 {
		fmt.Printf("Ошибка парсинга индекса бита: индекс должен быть положительным числом\n")
		os.Exit(1)
	}

	bitValue, err := strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil || (bitValue != 0 && bitValue != 1) {
		fmt.Printf("Ошибка парсинга значения бита: должно быть 0 или 1\n")
		os.Exit(1)
	}

	mask := int64(1) << (bitIndex - 1) // Маска с 1 в нужном месте

	var result int64
	if bitValue == 1 {
		result = num | mask // Установка бита в 1
	} else {
		result = num &^ mask // Установка бита в 0
	}

	fmt.Printf("Исходное число: %d (%b)\n", num, num)
	fmt.Printf("Установка %d-го бита в %d: %d (%b)\n", bitIndex, bitValue, result, result)
}
