package wallet

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

type ApplePay struct {
	Money int
	AppleID string
}

func (a *ApplePay) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("Не хватает денег на аккаунте")
	}
	a.Money -= amount
	return nil
}

// ------------------
// Ну и также наш интерфейс "Плательщик" (Payer),
// который требует только, чтобы был метод Pay
type Payer interface {
	Pay(int) error
}

// ------------------
// И функция "Купить" (Buy)
func Buy(p Payer) {
	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Ошибка при оплате %T: %v\n\n", p, err)
		return
	}
	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}

// ------------------