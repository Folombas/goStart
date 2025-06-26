package variables

import "fmt"

func main() {
	// int - платформозависимый тип, 32/64
	var i int = 10

	// автоматически выбранный int
	var autoInt = -10

	// int8, int16, int32, int64
	var bigInt int64 = 1<<32 - 1

	// платформозависимый тип, 32/64
	var unsignedInt uint = 100500

	// uint8, uint16, uint32, uint64
	var unsignedBigInt uint64 = 1<<64 - 1

	fmt.Println(i, autoInt, bigInt, unsignedInt, unsignedBigInt)
}