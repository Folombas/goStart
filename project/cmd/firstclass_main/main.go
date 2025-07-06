package main

import (
	"fmt"
	"project/data_types/firstclass"
)

func main() {
	fmt.Println("🚀 Запуск демонстрации работы с функциями как объектами первого класса")
	fmt.Println("========================================")

	firstclass.DoNothing()
	firstclass.FuncAnonymous()
	firstclass.CallbackFunc()

	fmt.Println("========================================")
	fmt.Println("🎉 Программа успешно завершена!")
}
