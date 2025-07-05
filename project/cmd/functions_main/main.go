package main

import (
	"fmt"
	"project/data_types/functions"
)

func main() {
	fmt.Println("🚀 Запуск программы ")
	fmt.Println("========================================")

	var multPrint int = functions.MultIn(5, 6, 2)
	println(multPrint)

	fmt.Println("========================================")
	fmt.Println("🎉 Программа успешно завершена!")
}
