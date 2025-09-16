// go run main.go
package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// по условию (сама завершится)
func workerWithCondition(stopAfter int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < stopAfter; i++ {
		fmt.Printf("Выполнение горутины, итерация %d\n", i)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("Горутина завершила работу по условию")
}

// канал уведомления
func workerWithChannel(stopChan <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stopChan:
			fmt.Println("Получен сигнал остановки через канал")
			return
		default:
			fmt.Println("Выполнение горутины...")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

// контекст
func workerWithContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Контекст отменен:", ctx.Err())
			return
		default:
			fmt.Println("Ожидание отмены контекста...")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

// runtime.Goexit()
func workerWithGoexit(wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Горутина завершена через Goexit()")

	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("Вызов Goexit()")
			runtime.Goexit()
		}
		fmt.Printf("Ожидание Goexit: %d\n", i)
		time.Sleep(200 * time.Millisecond)
	}
}

// канал с таймаутом
func workerWithTimeout(timeout time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	timeoutChan := time.After(timeout)

	for {
		select {
		case <-timeoutChan:
			fmt.Println("Время работы истекло")
			return
		default:
			fmt.Println("Ожидание таймаута...")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

// паника и восстановление
func workerWithPanic(wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Восстановлено после паники:", r)
		}
	}()

	for i := 0; i < 10; i++ {
		if i == 4 {
			panic("критическая ошибка!")
		}
		fmt.Printf("Ждём паники. %d\n", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("1. Завершение по условию:")
	wg.Add(1)
	go workerWithCondition(3, &wg)
	wg.Wait()
	fmt.Println()

	fmt.Println("2. Завершение через канал:")
	stopChan := make(chan struct{})
	wg.Add(1)
	go workerWithChannel(stopChan, &wg)
	time.Sleep(800 * time.Millisecond)
	close(stopChan)
	wg.Wait()
	fmt.Println()

	fmt.Println("3. Завершение через контекст:")
	ctx, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)
	defer cancel()
	wg.Add(1)
	go workerWithContext(ctx, &wg)
	wg.Wait()
	fmt.Println()

	fmt.Println("4. Завершение через runtime.Goexit():")
	wg.Add(1)
	go workerWithGoexit(&wg)
	wg.Wait()
	fmt.Println()

	fmt.Println("5. Завершение по таймауту:")
	wg.Add(1)
	go workerWithTimeout(800*time.Millisecond, &wg)
	wg.Wait()
	fmt.Println()

	fmt.Println("6. Завершение через панику:")
	wg.Add(1)
	go workerWithPanic(&wg)
	wg.Wait()
	fmt.Println()
}
