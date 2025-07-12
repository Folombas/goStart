package basic

import "fmt"

// ------------------
// Теперь рассмотрим несколько более сложный пример, когда
// у нас есть несколько структур, которые реализуют интерфейс
type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Не хватает денег в кошельке")
	}
	w.Cash -= amount
	return nil
}
// У нас тут есть кошелёк (Wallrt). Он реализует метод Pay. 
// У него есть какое-то количество денег в нём.

// ------------------
// У нас есть карточка с некоторой структурой.
// И она тоже реализует метод Pay
type Card struct {
	Balance int
	ValidUntil string
	Cardholder string
	CVV        string
	Number     string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("Не хватает денег на карте")
	}
	c.Balance -= amount
	return nil
}

// ------------------
// И есть ещё Apple Pay

type AplplePay struct {
	Money int
	AppleID string
}

func (a *AplplePay) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("Не хватает денег на аккаунте")
	}
	a.Money -= amount
	return nil
}

// ------------------
// Ну и также наш интерфейс "Плательщик" (Payer)
type Payer interface {
	Pay(int) error
}

// ------------------