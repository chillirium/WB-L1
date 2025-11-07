package main

import (
	"fmt"
	"reflect"
	"sync"
)

func checkType(v any) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		if reflect.TypeOf(v).Kind() == reflect.Chan {
			fmt.Printf("%T\n", v)
		} else {
			fmt.Println("unknown")
		}
	}
}

func main() {
	inp := make(chan any)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for v := range inp {
			checkType(v)
		}
	}()

	inp <- 100
	inp <- "Go"
	inp <- false
	inp <- make(chan struct{})
	inp <- 1.

	close(inp)
	wg.Wait()
}
