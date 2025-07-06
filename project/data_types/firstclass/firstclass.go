package firstclass

import "fmt"

// Обычная функция
func DoNothing() {
	fmt.Println("I`m regular function")
}

func FuncAnonymous() {
	// Анонимная функция
	func(in string) {
		fmt.Println("Anon func out:", in)
	}("nobody	")
}

func CallbackFunc() {
	// присваивание анонимной функции в переменную
	printer := func(in string) {
		fmt.Println("printer outs:", in)
	}
	printer("as variable")


// Определяем тип функции
type strFuncType func(string)

// Функция принимает коллбек
worker := func(callback strFuncType) {
	callback("as callback")
}
worker(printer)
}
