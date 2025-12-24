package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите строку: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}

	input = strings.TrimSpace(input)

	result := reverseString(input)
	fmt.Println("Результат:", result)
}
