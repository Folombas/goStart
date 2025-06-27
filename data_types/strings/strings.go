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

	helloWord := "Привет Мир"
	// конкатенация строк
	andGoodMorning := helloWord + " и доброе утро!"

	// строки неизменяемы
	// cannot assign to helloWord[0]
	helloWord[0] = 72

	// получение длины строки
	byteLen := len(helloWord)						// 19 bytes
	symbols := utf8.RuneCountInString(helloWord) 	// 12 runes

	// получение подстроки, в байтах, не символах!
	hello := helloWorld[:12] // Привет, 0-11 байты
	H := helloWorld[0]		// byte, 72, не "П"	

	// конвертация в слайс байт и обратно
	byteString = []byte(helloWorld)
	helloWorld = string(byteString)


}
