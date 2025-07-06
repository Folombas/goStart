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

// // Присваивание анонимной функции в переменную
// printer := func(in string) {
// 	fmt.Println("printer outs:", in)
// }
// printer("as variable")
