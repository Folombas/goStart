package main

import (
	"fmt"
	"project/data_types/slices"
)

func main() {
	fmt.Println("🚀 Запуск полной демонстрации слайсов")
	fmt.Println("========================================")

	fmt.Println("📝 Это Демонстрация слайсов")

	slices.Demo_Slices() // Функция из slices.go

	fmt.Println("Копирование слайсов")

	slices.Demo_Slices2() // Функция из slices_2.go

	fmt.Println("========================================")
	fmt.Println("🎉 Программа успешно завершена!")
}
