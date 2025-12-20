package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Mutex
type MutexCounter struct {
	mu    sync.Mutex
	value int64
}

func (c *MutexCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *MutexCounter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// Atomic
type AtomicCounter struct {
	value atomic.Int64
}

func (c *AtomicCounter) Increment() {
	c.value.Add(1)
}

func (c *AtomicCounter) Value() int64 {
	return c.value.Load()
}

// интерфейс счётчика
type Counter interface {
	Increment()
	Value() int64
}

func testCounter(name string,
	counter Counter,
	numGoroutines,
	incrementsPerGoroutine int) time.Duration {

	var wg sync.WaitGroup

	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)
	count := counter.Value()

	fmt.Printf("%s итого: %d\n", name, count)

	return elapsed
}

func main() {
	numGoroutines := 200
	incrementsPerGoroutine := 300

	fmt.Printf("Горутин: %d\n", numGoroutines)
	fmt.Printf("Инкрементов на горутину: %d\n", incrementsPerGoroutine)
	fmt.Printf("Ожидаемое значение: %d\n\n", numGoroutines*incrementsPerGoroutine)

	mutexCounter := &MutexCounter{}
	timeMutex := testCounter("sync.Mutex", mutexCounter, numGoroutines, incrementsPerGoroutine)
	fmt.Printf("Время выполнения: %v\n\n", timeMutex.Round(time.Microsecond))

	atomicCounter := &AtomicCounter{}
	timeAtomic := testCounter("sync/atomic", atomicCounter, numGoroutines, incrementsPerGoroutine)
	fmt.Printf("Время выполнения: %v\n\n", timeAtomic.Round(time.Microsecond))
}
