package firstclass

import (
	"fmt"
	"strings"
)

// 1. Обычная функция с именем
func DiagnoseSasha(symptoms string) string {
	if strings.Contains(symptoms, "велосипед") {
		return "Хронический велокурьеризм 2-й степени"
	}
	return "Ипохондрия программиста"
}

func FuncVar() {
	//Присваиваем функцию переменной
	transport := func(distance float64) string {
		if distance > 5 {
			return "Велосипед"
		}
		return "Пешком ножками"
	}

	fmt.Println("Сегодня:", transport(3))
}

// Функция, принимающая колбэк
func CookOmelette(ingredients string, callback func(result string)) {
	fmt.Println("Секретарша готовит омлет с:", ingredients)
	callback("Идеальный омлет с 2 яйцами и сарделькой")
}

// Функция с замыканием
func EarningsCounter() func(float64) float64 {
	total := 0.0 // Переменная захватывается замыканием
	return func(earned float64) float64 {
		total += earned
		fmt.Printf("Заработано сегодня: %.2f руб. Всего: %.2f руб.\n", earned, total)
		return total
	}
}