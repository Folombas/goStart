package slices

import "fmt"

func Demo_Slices2() {
	buf := []int{1, 2, 3, 4, 5}
	fmt.Println(buf)

	// Получение среза, указывающего на ту же память
	sl1 := buf[1:4] // [2, 3, 4]
	sl2 := buf[:2] // [1, 2]
	sl3 := buf[2:] // [3, 4, 5]
	fmt.Println(sl1, sl2, sl3)

	newBuf := buf[:] // [1, 2, 3, 4, 5]
	// buf = [9, 2, 3, 4, 5], так как newBuf указывает на buf, то есть на ту же память
	newBuf[0] = 9

	// newBuf теперь указывает на другие данные
	newBuf = append(newBuf, 6)

	// buf = [9, 2, 3, 4, 5], не изменился
	// newBuf = [1, 2, 3, 4, 5, 6], изменился
	newBuf[0] = 1
	fmt.Println("buf", buf)
	fmt.Println("newBuf", newBuf)

	// Копирование одного слайса в другой
	var emptyBuf []int // len=0, cap=0
	// неправильно - скопирует меньшее (по len) из 2-х слайсов
	copied := copy(emptyBuf, buf) // copied = 0
	fmt.Println(copied, emptyBuf)

	// Правильно
	newBuf = make([]int, len(buf), cap(buf))
	copy(newBuf, buf) // newBuf = [1, 2, 3, 4, 5]
	fmt.Println(newBuf)

	// Можно копировать в часть существующего слайса
	ints := []int{1, 2, 3, 4}
	copy(ints[1:3], []int{5, 6}) // ints = [1, 5, 6, 4]
	fmt.Println(ints)
}