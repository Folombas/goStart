package slices

import "fmt"

// Demp_Slices - это экспортируемая функция для демонстрации работы с слайсами
func Demo_Slices() {
	// создание
	var buf0 []int				// len=0, cap=0
	buf1 := []int{}				// len=0, cap=0
	buf2 := []int{42}			// len=1, cap=1
	buf3 := make([]int, 0)  	// len=0, cap=0
	buf4 := make([]int, 5)		// len=5, cap=5
	buf5 := make([]int, 5, 10)  // len=5, cap=10

	fmt.Println("Создание срезов:")
	fmt.Println("buf0:", buf0, "len:", len(buf0), "cap:", cap(buf0))
	fmt.Println("buf1:", buf1, "len:", len(buf1), "cap:", cap(buf1))
	fmt.Println("buf2:", buf2, "len:", len(buf2), "cap:", cap(buf2))
	fmt.Println("buf3:", buf3, "len:", len(buf3), "cap:", cap(buf3))
	fmt.Println("buf4:", buf4, "len:", len(buf4), "cap:", cap(buf4))
	fmt.Println("buf5:", buf5, "len:", len(buf5), "cap:", cap(buf5))

	// обращение к элементам
	fmt.Println("\nОбращение к элементам:")
	someInt := buf2[0]
	fmt.Println("buf2[0]: ", someInt)

	// ошибка при выполнении
	// panic: runtime error: index out of range
	// someOtherInt := buf2[1]
}