package strings

import (
	"unicode/utf8"
)

func main() {
	// пустая строка по-умолчанию
	var str string

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

	helloWord := "Привет, мир!"
	// конкатенация строк
	andGoodMorning := helloWord + " и доброе утро!"

	// получение длины строки
	byteLen := len(helloWord)
	symbols := utf8.RuneCountInString(helloWord)

}
