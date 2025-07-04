package functions

import "fmt"

// обычное объявление функции
func singleIn(in int) int {
	return in
}

// Много параметров
func multIn(a, b int, c int) int {
	return a + b + c
}

// Именованный результат
func namedReturn() (out int) {
	out = 3
	return
}

// Несколько результатов
func multiplrRerutn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("some error happend")
	}
	return in, nil
}