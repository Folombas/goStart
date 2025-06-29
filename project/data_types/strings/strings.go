package strings

import (
	"fmt"
	"unicode/utf8"
)

func DemoStrings() {
	// пустая строка по-умолчанию
	var str string
	str = "This is a string"

	// со спец символами
	var hello string = "Привет\n\t"

	// без спец символов
	var world string = `Мир\n\t`

	// UTF-8 из коробки
	var helloWord = "Привет, мир!"
	hi := "你好, 世界！"

	// одинарные кавычки для байт (uint8)
	var rawBinary byte = '\x27'

	// rune (uint32) для UTF-8 символов
	var someChinese rune = '好'

	helloWord = "Привет Мир"
	// конкатенация строк
	andGoodMorning := helloWord + " и доброе утро!"

	// строки неизменяемы
	// cannot assign to helloWord[0]
	// helloWord[0] := 72

	// получение длины строки
	byteLen := len(helloWord)						// 19 bytes
	symbols := utf8.RuneCountInString(helloWord) 	// 12 runes

	fmt.Println(str)
	fmt.Println(hello)
	fmt.Println(world)
	fmt.Println(hi)
	fmt.Println(rawBinary)
	fmt.Println(someChinese)
	fmt.Println(andGoodMorning)
	fmt.Println(byteLen)
	fmt.Println(symbols)
	fmt.Println(hello)
	fmt.Println("Начинаем ежедневное изучение языка программирования Go!")
}
