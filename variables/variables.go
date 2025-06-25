package variables

import "fmt"

func main() {
	// значение по умолчанию
	var num0 int

	// значение при инициализации
	var num1 int = 1

	// пропуск типа
	var num2 = 20
	fmt.Println(num0, num1, num2)

	// короткое объявление переменной
	num := 30
	// только для новых переменных
	// no new variables on left side of :=
	// num := 31

	num += 1
}
