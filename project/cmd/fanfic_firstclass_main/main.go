package main

import (
	"fmt"
	"project/data_types/firstclass"
)

func main() {
	fmt.Println("🚀 Запуск демонстрации работы с функциями как объектами первого класса")
	fmt.Println("========================================")

	diagnosis := firstclass.DiagnoseSasha("Жара, велик, 497 рублей")
	fmt.Println("Токсиков: Ваш диагноз -", diagnosis)
	fmt.Println("----------------------------------------")
	// Анонимная функция без имени (Барбоскин покупает айс-латте)
	buyIceLatte := func(price float64) {
		fmt.Printf("Барбоскин тратит %.2f рублей на айс-латте\n", price)
		fmt.Println("Саша: Да это же моя дневная зарплата за доставку!")
	}

	buyIceLatte(750) // Вызов анонимной функции

	firstclass.FuncVar()

	// Передаём функцию в качестве аргумента
	firstclass.CookOmelette("Яйца, сарделька", func(result string) {
		fmt.Println("Токсиков: *нюхает еду*", result)
		fmt.Println("Саша: А у меня рисовая каша...")
	})

	
		counter := firstclass.EarningsCounter()
		counter(497.27) // Доставка 1
		counter(320.50) // Доставка 2
	

	fmt.Println("========================================")
	fmt.Println("🎉 Программа успешно завершена!")
}

