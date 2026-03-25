package main

import (
	"fmt"
	"project/data_types/basic"
)

func main() {
	// Выводим заголовок
	basic.PrintHeader("Первая программа на Go - Hello Go!")

	// Вызываем основную функцию
	basic.HelloGo()

	// Дополнительная практика - объявление переменных
	fmt.Println("📝 Практика объявления переменных:")
	fmt.Println()

	// Способ 1: var с типом
	var age int = 25
	fmt.Printf("  var age int = %d\n", age)

	// Способ 2: var без типа
	var name = "Alex"
	fmt.Printf("  var name = \"%s\"\n", name)

	// Способ 3: короткое объявление :=
	city := "Moscow"
	fmt.Printf("  city := \"%s\"\n", city)

	// Способ 4: множественное присваивание
	var x, y = 10, 20
	fmt.Printf("  var x, y = %d, %d\n", x, y)

	// Способ 5: множественное короткое
	a, b := 5, 15
	fmt.Printf("  a, b := %d, %d\n", a, b)
	fmt.Println()

	// Пример с константами
	const Pi = 3.14159
	const MaxUsers = 100
	fmt.Println("📌 Константы:")
	fmt.Printf("  Pi = %.2f\n", Pi)
	fmt.Printf("  MaxUsers = %d\n", MaxUsers)
	fmt.Println()

	// Арифметические операции
	fmt.Println("🔢 Арифметические операции:")
	num1, num2 := 42, 8
	fmt.Printf("  %d + %d = %d\n", num1, num2, num1+num2)
	fmt.Printf("  %d - %d = %d\n", num1, num2, num1-num2)
	fmt.Printf("  %d * %d = %d\n", num1, num2, num1*num2)
	fmt.Printf("  %d / %d = %d\n", num1, num2, num1/num2)
	fmt.Printf("  %d %% %d = %d\n", num1, num2, num1%num2)
	fmt.Println()

	// Булевы операции
	fmt.Println("🔍 Булевы операции:")
	isTrue, isFalse := true, false
	fmt.Printf("  true && false = %v\n", isTrue && isFalse)
	fmt.Printf("  true || false = %v\n", isTrue || isFalse)
	fmt.Printf("  !true = %v\n", !isTrue)
	fmt.Println()

	// Строковые операции
	fmt.Println("📝 Строковые операции:")
	first := "Hello"
	second := "Go"
	combined := first + ", " + second + "!"
	fmt.Printf("  \"%s\" + \", \" + \"%s\" + \"!\" = \"%s\"\n", first, second, combined)
	fmt.Println()

	// Выводим подвал
	basic.PrintFooter()
}
