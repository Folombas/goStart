package structs

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	Id      int
	Name    string
	Cleaner func(string) string
	Owner   Person
	Person
}

func DemoStructs() {
	// Полное объявление структуры
	var acc Account = Account{
		Id:   1,
		Name: "Gosha",
		Person: Person{
			Name:	  "Гоша",
			Address: "Москва",
		},
	}
	fmt.Printf("%#v\n", acc)

	// короткое объявление структуры
	acc.Owner = Person{2, "Gosha Golang", "Moscow"}

	fmt.Printf("%#v\n", acc)

	fmt.Println(acc.Name)
	fmt.Println(acc.Person.Name)
}
