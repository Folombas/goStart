package main

import (
	"fmt"
	"project/data_types/interface_composition"
)

func main() {
	fmt.Println("🚀 Запуск демонстрации работы композиций интерфейсов")
	fmt.Println("========================================")

	myPhone := interface_composition.Phone{Money: 9}
	interface_composition.PayMetroWithPhone(&myPhone)

	fmt.Println("========================================")
	fmt.Println("🎉 Программа успешно завершена!")
}
