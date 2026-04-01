package main

import (
	"fmt"
	"project/data_types/errors_pkg"
)

func main() {
	fmt.Println("🚀 Запуск демонстрации обработки ошибок")
	fmt.Println("========================================")

	errors_pkg.DemoErrors()

	fmt.Println("========================================")
	fmt.Println("🎉 Программа успешно завершена!")
}
