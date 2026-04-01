package main

import (
	"fmt"
	"project/data_types/channels"
)

func main() {
	fmt.Println("🚀 Запуск демонстрации каналов и горутин")
	fmt.Println("========================================")

	channels.DemoChannels()

	fmt.Println("========================================")
	fmt.Println("🎉 Программа успешно завершена!")
}
