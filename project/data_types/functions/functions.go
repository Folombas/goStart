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
	out = 2
	return
}

// Несколько результатов
func multipleReturn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("some error happend")
	}
	return in, nil
}

// несколько именованных результатов
func multipleNamedReturn(ok bool) (rez int, err error) {
	rez = 1
	if ok {
		err = fmt.Errorf("some error happend")
		// аналогично return rez, err
		// или return 1, fmt.Errorf("some error happend")
		return
	}
	rez = 2
	return
}

// не фиксированное количество  параметров
func sum(in ...int) (result int) {
	fmt.Printf("in := %#v \n", in)
	for _, val := range in {
		result += val
	}
	return
}
