package main

import "fmt"

func createSet(items []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, item := range items {
		set[item] = struct{}{}
	}
	return set
}

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	set := createSet(words)

	setItems := make([]string, 0, len(set))

	for word := range set {
		setItems = append(setItems, word)
	}

	fmt.Println(setItems)
}
