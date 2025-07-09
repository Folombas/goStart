package structs

import "fmt"

// Courier представляет сущность курьера
type Courier struct {
	Name       string
	Transport  string
	Depressed  bool
	Earnings   float64
}

func FanficStructs() {
	// Создаём героя-курьера
	sasha := Courier{
		Name:      "Саша",
		Transport: "Велосипед",
		Depressed: true,
		Earnings:  497.28,
	}

	fmt.Printf("%s заработал: %.2f руб.\n", sasha.Name, sasha.Earnings)
}