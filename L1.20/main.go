package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseRunes(runes []rune, start, end int) {
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
}

func reverseWords(s string) string {
	runes := []rune(s)
	n := len(runes)

	if n == 0 {
		return ""
	}

	// "snow dog sun" → "nus god wons"
	reverseRunes(runes, 0, n-1)

	// "nus god wons" → "sun dog snow"
	wordStart := 0

	for i := 0; i <= n; i++ {
		if i == n || runes[i] == ' ' {
			reverseRunes(runes, wordStart, i-1)
			wordStart = i + 1
		}
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

	input = strings.TrimRight(input, "\r\n")
	result := reverseWords(input)

	fmt.Println("Результат:", result)
}
