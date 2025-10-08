// "Пример: -20:{-25.4, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20:{24.5}, 30:{32.5}."
// группа 0 тогда от -10 до 10 не включая края
// go run main.go
package main

import (
	"fmt"
	"sort"
	"strings"
)

// форматирорвание как в примере, но не всё в одну строку, точность до одного знакаа после запятой
func formatFloats(vals []float64) string {
	strs := make([]string, len(vals))
	for i, v := range vals {
		strs[i] = fmt.Sprintf("%.1f", v)
	}
	return "{" + strings.Join(strs, ", ") + "}"
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	//если есть повторения, то они останутся
	for _, temp := range temperatures {
		key := int(temp) / 10 * 10
		groups[key] = append(groups[key], temp)
	}

	//"Порядок в подмножествах не важен", но наверно важен порядок групп
	keys := make([]int, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, key := range keys {
		fmt.Printf("%d:%s\n", key, formatFloats(groups[key]))
	}
}
