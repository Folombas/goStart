package map_learn

import "fmt"

func CourierMap() {
	// 1. Создаём карту для хранения заказов (ключ: ID заказа, значение: адрес доставки)
	pendingOrders := make(map[int]string)
	fmt.Println("Карта создана:", pendingOrders)

	// 2. Добавляем заказы как в курьерском приложении
	pendingOrders[101] = "Химки, ул. Гагарина, 10"
	pendingOrders[202] = "Москва, Ленинский пр-т, 100"
	pendingOrders[303] = "Лобня, ул. Пушкина, 12"
	fmt.Println("\nПосле добавления заказов:")
	printOrders(pendingOrders)

	// 3. Проверяем существование заказа (как проверка перед выездом)
	orderID := 202
	if address, exists := pendingOrders[orderID]; exists {
		fmt.Printf("\nЗаказ %d существует! Адрес: %s\n", orderID, address)
	} else {
		fmt.Printf("\nЗаказ %d не найден!\n", orderID)
	}

	// 4. Обновляем адрес (клиент изменил адрес)
	fmt.Println("\nКлиент изменил адрес для заказа 101...")
	pendingOrders[101] = "Химки, Московский пр-т, 15"
	fmt.Println("Обновлённый заказ 101:", pendingOrders[101])

	// 5. Доставляем заказ (удаляем из системы)
	fmt.Println("\nДоставляем заказ 202...")
	delete(pendingOrders, 202)
	fmt.Println("После доставки:")
	printOrders(pendingOrders)

	// 6. Попытка доставить несуществующий заказ
	fmt.Println("\nПробуем доставить несуществующий заказ 404...")
	if _, exists := pendingOrders[404]; !exists {
		fmt.Println("ОШИБКА: Заказ 404 не найден! Звоним в службу поддержку Диме")
	}

	// 7. Чай с мамой (добавляем несколько заказов сразу)
	fmt.Println("\nВечер: попил на кухне чай с мамой, нашёл новые заказы в мобильных курьерских приложениях...")
	newOrders := map[int]string{
		404: "Путилово, ул. Космонавтов, 1",
		505: "Дедовск, переулок Строителей, 10",
	}
	for id, address := range newOrders {
		pendingOrders[id] = address
	}
	fmt.Println("Все заказы после домашнего вечернего чая с мамой:")
	printOrders(pendingOrders)

	// 8. Статистика за день
	fmt.Printf("\nИтоги дня:\nВсего заказов: %d\n", len(pendingOrders))
}

// Вспомогательная функция для печати заказов
func printOrders(orders map[int]string) {
	fmt.Println("Текущие заказы:")
	fmt.Println("ID\tАдрес")
	fmt.Println("--------------------------------------")
	for id, address := range orders {
		fmt.Printf("%d\t%s\n", id, address)
	}
}
