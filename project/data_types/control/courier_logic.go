package control

import (
	"fmt"
	"math/rand"
)

func CourierLogic(){

	// 1. Простое условие if Проверка погоды
	isRaining := rand.Float32() < 0.7 // 70% chance of rain
	fmt.Println("=== Погодный инспектор ===")
	if isRaining {
		fmt.Println("☔️ Нужен дождевик! Сегодня мокрый день")
	} else {
		fmt.Println("☀️ Отличный день для велопрогулок!")
	}

	// 2. Условие с блоком инициализации - Проверка заказа
	fmt.Println("\n=== Проверка заказа ===")
	if weight := rand.Intn(15) + 1; weight > 10 {
		fmt.Printf("⚠️ Внимание! Тяжелый груз (%d кг). Возьмите рюкзак\n", weight)
	} else if weight > 5 {
		fmt.Printf("🛒 Средний вес (%d кг). Курьерская сумка подойдет\n", weight)
	} else {
		fmt.Printf("📦 Легкий заказ (%d кг). Можно взять в ручную\n", weight)
	}

	// 3. Множественный if else - Обработка проблем
	fmt.Println("\n=== Ситуации на маршруте ===")
	clientAvailable := rand.Intn(100) > 30 // 70% chance available
	addressCorrect := rand.Intn(100) > 20 // 80% chance correct
	paymentReady := rand.Intn(100) > 10 // 90% chance ready

	if !addressCorrect && !clientAvailable {
		fmt.Println("🚨 КРИТИЧЕСКАЯ СИТУАЦИЯ: Адрес неверный и клиент недоступен!")
	} else if !addressCorrect {
		fmt.Println("⚠️ Проблема: Адрес доставки указан неверно")
	} else if !clientAvailable {
		fmt.Println("📞 Клиент не отвечает на звонки...")
	} else if !paymentReady {
		fmt.Println("💳 Проблема с оплатой заказа")
	} else {
		fmt.Println("✅ Все системы в норме. Приятной доставки!")
	}

	// 4. Switch по 1 переменной - Типа заказа
	fmt.Println("\n=== Типа заказа ===")
	orderTypes := []string{"еда", "аптека", "цветы", "документы", "электроника"}
	orderType := orderTypes[rand.Intn(len(orderTypes))]

	switch orderType {
	case "еда":
		fmt.Println("🍔 Заказ: Пищевые продукты. Требуется термосумка!")
	case "аптека":
		fmt.Println("💊 Заказ: Лекарства. Срочная доставка!")
	case "цветы":
		fmt.Println("💐 Заказ: Букет цветов. Аккуратно переворачивать!")
	case "документы":
		fmt.Println("📄 Заказ: Важные документы. Водонепроницаемая упаковка!")
	case "электроника":
		fmt.Println("📱 Заказ: Гаджеты. Особая осторожность!")
	default:
		fmt.Println("📦 Заказ: Стандартная посылка")
	}

	
}

