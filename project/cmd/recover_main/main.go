package main

import (
	"fmt"
	"project/data_types/recover"
)

func main() {
	fmt.Println("🚀 Запуск демонстрации работы с паникой")
	fmt.Println("========================================")

	// recover.DeferTest()
	recover.ProcessDelivery("VIP-777")

	fmt.Println("========================================")
	fmt.Println("🎉 Программа успешно завершена!")
}