package functions

import "fmt"

// обычное объявление функции
func SingleIn(in int) int {
	return in
}

// Много параметров
func MultIn(a, b int, c int) int {
	return a + b + c
}

// Именованный результат
func NamedReturn() (out int) {
	out = 2
	return
}

// Несколько результатов
func MultipleReturn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("some error happend")
	}
	return in, nil
}

// несколько именованных результатов
func MultipleNamedReturn(ok bool) (rez int, err error) {
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
func Sum(in ...int) (result int) {
	fmt.Printf("in := %#v \n", in)
	for _, val := range in {
		result += val
	}
	return
}
