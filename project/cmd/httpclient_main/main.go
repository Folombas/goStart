package main

import (
	"fmt"
	"project/data_types/httpclient"
)

func main() {
	fmt.Println("🚀 Запуск демонстрации HTTP Client")
	fmt.Println("========================================")

	httpclient.DemoHTTPClient()

	fmt.Println("========================================")
	fmt.Println("🎉 Программа успешно завершена!")
}
