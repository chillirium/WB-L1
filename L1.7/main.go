// go run -race main.go
package main

import (
	"fmt"
	"sync"
)

// map со встроенным sync.Mutex
type SafeMap struct {
	mu   sync.Mutex
	data map[string]int
}

// создание нового экземпляра экземпляр
func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

// запись
func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

// чтение
func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	val, ok := sm.data[key]
	return val, ok
}

// получить копию всей map (для безопасного вывода/проверки)
func (sm *SafeMap) GetAll() map[string]int {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// копия, чтобы не держать блокировку при итерации
	result := make(map[string]int, len(sm.data))
	for k, v := range sm.data {
		result[k] = v
	}
	return result
}

func main() {
	safeMap := NewSafeMap()
	var wg sync.WaitGroup

	// записываем 100 ключей
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.Set(fmt.Sprintf("key%d", i), i)
		}(i)
	}

	wg.Wait()

	allData := safeMap.GetAll()

	for key, value := range allData {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// проверка корректности
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%d", i)
		if val, ok := allData[key]; !ok || val != i {
			fmt.Printf("Ошибка: ключ %s не найден или значение неверное (получено: %v)\n", key, val)
		}
	}
}
