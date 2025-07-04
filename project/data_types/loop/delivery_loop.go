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
			fmt.Println("⏰ Рабочий день окончен! Курьер устал.")
			break
		}
	}

	// 2. Цикл с одиночным условием (while-подобный)
	fmt.Println("=== Поиск заказов в приложении ===")
	orderAvailable := false
	attempts := 0
	for !orderAvailable {
		attempts++
		fmt.Printf("Попытка #%d: поиск доступных заказов...\n", attempts)

		// Эмуляция появления заказа на 3-й попытке
		if attempts >= 3 {
			orderAvailable = true
			fmt.Println("🎉 Найден заказ! Адрес: Химки, ул. Центральная, 15")
		}
	}

	// 3. Цикл с инициализацией и условием (классический for)
	fmt.Println("=== Обработка заказов на складе ===")
	fmt.Println("Курьер проверяет заказы в сумке:")
	for i := 1; i <= 4; i++ {
		fmt.Printf("Заказ #%d: Проверен, упакован, готов к доставке", i)
	}
	fmt.Println("✅ Все заказы готовы к отправке!")

	// 4. Операции по slice (традиционный способ)
	fmt.Println("=== Планирование маршрута (традиционный способ) ===")
	addresses := []string{
		"ул. Ленина, 10",
		"пр-т Мира, 25",
		"ул. Центральная, 15",
		"б-р Космонавтов, 7",
	}

	fmt.Println("Оптимальный маршрут доставки:")
	for i := 0; i < len(addresses); i++ {
		fmt.Printf("%d. %s\n", i+1, addresses[i])
	}
	fmt.Println()

	// 5. Итерирование по slice с range (современный способ)
	// for idx, value := range slice {}
	fmt.Println("=== Доставка заказов (range для slice) ===")
	packages := []struct {
		ID        int
		Address   string
		Weight    float64
	}{
		{101, "ул. Садовая, 42", 2.5},
		{202, "пр-т Победы, 17", 1.8},
		{303, "ул. Лесная, 5", 3.2},
	}

	fmt.Println("Начало доставки:")
	for idx, pkg := range packages {
		fmt.Printf("Доставка #%d: Заказ %d (%.1f кг) -> %s\n",
			idx+1, pkg.ID, pkg.Weight, pkg.Address)
	}
	fmt.Println("✅ Все заказы доставлены!")

	// 6. Итерирование по map с range
	fmt.Println("=== Статистика дня (range для map) ===")
	deliveryStats := map[string]int{
		"успешно":   8,
		"отменено":  2,
		"возврат":   1,
		"опозданий": 1,
	}
	
	total := 0
	for status, count := range deliveryStats {
		fmt.Printf("%s: %d\n", status, count)
		total += count
	}
	fmt.Printf("Всего обработано заказов: %d\n\n", total)

	// 7. Итерирование по строке с range
	// for idx, runeValue := range str {}
	fmt.Println("=== Обработка адреса (range для строки) ===")
	address := "ул. Парковая, 27к2"

	// 7.1 Неправильный подход: итерация по байтам
	fmt.Println("Адрес как байты:", []byte(address))
	fmt.Print("Символы (по байтам): ")
	for i := 0; i < len(address); i++ {
		fmt.Printf("%c ", address[i])
	}
	fmt.Println("\n⚠️ Проблема: кириллица и символ 'к' обрабатываются некорректно")

	// 7.2 Правильный подход: итерация по рунам
	fmt.Print("\nСимволы (по рунам): ")
	for idx, runeValue := range address {
		fmt.Printf("%c(%d) ", runeValue, idx)
	}

	// 7.3 Подсчет символов 
	fmt.Printf("\n\nАнализ адреса: \"%s\"", address)
	fmt.Printf("\nДлина в байтах: %d", len(address))
	fmt.Printf("\nКоличество символов: %d", utf8.RuneCountInString(address))
	fmt.Printf("\nПервый символ: %c", []rune(address)[0])
}

