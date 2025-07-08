package recover

import "fmt"

func ProcessDelivery(orderID string) {
	fmt.Printf("🚴‍♂️ Начало доставки заказа %s\n", orderID)

	// Отложенные вызовы выполняются в обратном порядке (LIFO - Last In, First Out, означает - последний вошёл, первый вышел)
	defer fmt.Println("📦 Заказ доставлен (обязательное действие)")
	defer fmt.Println("📱 Отметка в мобильном приложении")

	if orderID == "VIP-777" {
		defer fmt.Println("✨ Особый статус: Доставка доктору Лесли Токсикову!")
	}

	fmt.Println("🛣️ Маршрут построен...")
	fmt.Println("🚲 Поездка началась...")
}