// go run main.go <кол-во воркеров>
// Код тот же, что и в прошлой задаче. Обоснование подхода:
// проще и нагляднее контекста в учебной задаче,
// отсутствуют накладные расходы за дополнительную функциональность.
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// Воркер читает канал, пишет в стандартный вывод
func worker(id int, dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range dataChan {
		fmt.Printf("Воркер %d: %d\n", id, num)
	}
	fmt.Printf("Воркер %d завершил работу\n", id)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Необходимо указать только количество воркеров")
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Ошибка, количество воркеров должно быть int: %v", err)
	}

	dataChan := make(chan int)
	var wg sync.WaitGroup

	// Создание воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChan, &wg)
	}

	// Обработка сигнала останова
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Главная горутина, пишет счётчик в канал
	go func() {
		counter := 0
		for {
			select {
			case <-sigChan:
				// Сигнал завершения (ctrl+C)
				close(dataChan)
				return
			default:
				dataChan <- counter
				counter++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
	fmt.Println("Все воркеры завершили работу. Программа завершена.")
}
