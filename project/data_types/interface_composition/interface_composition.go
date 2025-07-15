package interface_composition

import "fmt"




type Phone struct {
	Money   int
	AppleID string
}

func (p *Phone) Pay(amount int) error {
	if p.Money < amount {
		return fmt.Errorf("Not enough money on account")
	}
	p.Money -= amount
	return nil
}

func (p *Phone) Ring(number string) error {
	if number == "" {
		return fmt.Errorf("Please, enter phone")
	}
	return nil
}

// Встраивание интерфейсов

// Подобно структурамЮ когда мы можем вложить (заимбедить) одну структуру в другую,
// и иметь доступ к её полям,
// похожую операцию можно выполнить и с интерфейсами.

// Мы можем встраивать один интерфейс в другой,
// тем самым образуя более сложные интерфейсы.

// У нас тут ниже есть интерфейс "Платильщик",
// который требует метод "Pay", который принимает целое число и возвращает ошибку.
type Payer interface {
	Pay(int) error
}

// А тут ниже есть интерфейс "Звонилка" (Ringer),
// который требует метод "Позвонить" (Ring), который принимает строку и возвращает ошибку.
type Ringer interface {
	Ring(string) error
}

// И есть интерфейс "Смартфон с функицей NFC",
// который реализует метод бесконтактной оплаты через NFC-технологию.
// И этот интефейс образован композицией двух других интерфейсов - Платильщик (Payer) и Звонилка (Ringer).
type NFCPhone interface {
	Payer
	Ringer
}

// И получается, что наш интерфейс полностью выглядит так:
type NFCSmartphone interface {
	Payer(int) error
	Ringer(string) error
}

// Но каждый раз полностью объявлять новый тип - это не очень удобно,
// поэтому вот композиция интерфейсов позволяет нам облегчить работу Go-программиста.


func PayMetroWithPhone(phone NFCPhone) {
	err := phone.Pay(1)
	if err != nil {
		fmt.Printf("Ошибка при оплате: %v\n\n", err)
		return
	}
	fmt.Printf("Турникет открыт через %T\n\n", phone)
}