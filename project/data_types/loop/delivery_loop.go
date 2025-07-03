package loop

import (
	"fmt"
	"unicode/utf8"
)

func DeliveryLoop() {
	// 1. Бесконечный цикл без условия (имитация рабочего дгя курьера)
	fmt.Println("=== Начало рабочего дня ===")
	deliveryCount := 1
	for {
		fmt.Printf("Доставка #%d завершена\n", deliveryCount)
		deliveryCount++

		// Условие выхода: курьер устал после 5 доставок
		if deliveryCount > 5 {
			fmt.Println("⏰ Рабочий день окончен! Курьер устал.\n")
			break
		}
	}

	// 2. Цикл с одиночеым условием (while-подобный)
	fmt.Println("=== Поиск заказов в приложении ===")
	orderAvailable := false
	attempts := 0
	for !orderAvailable {
		attempts++
		fmt.Printf("Попытка #%d: поиск доступных заказов...\n", attempts)

		// Эмуляция появления заказа на 3-й попытке
		if attempts >= 3 {
			orderAvailable = true
			fmt.Println("🎉 Найден заказ! Адрес: Химки, ул. Центральная, 15\n")
		}
	}

	// 3. Цикл с инициализацией и условием (улассический for)
	fmt.Println("=== Обработка заказов на складе ===")
	fmt.Println("Курьер проверяет заказы в сумке:")
	for i := 1; i <= 4; i++ {
		fmt.Println("Заказ #%d: Проверен, упакован, готов к доставке\n", i)
	}
	fmt.Println(""✅ Все заказы готовы к отправке!\n"")

	// 4. Операции по slice (традиционный способ)

}

