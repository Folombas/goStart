package main

import (
	"fmt"
	"project/data_types/basic"
)

func main() {
	fmt.Println("🚀 Запуск демонстрации работы с интерфейсами")
	fmt.Println("========================================")

	myWallet := &basic.Wallet{Cash: 100}
	basic.Buy(myWallet)

	fmt.Println("========================================")
	fmt.Println("🎉 Программа успешно завершена!")
}
