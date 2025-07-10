package methods

import "fmt"

type Person struct {
	Id   int
	Name string
}

// Не изменит оригинальной структуры, для которой вызван
func (p Person) UpdateName(name string) {
	p.Name = name
}

// Изменяет оригинальную структуру
func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

func (p *Account) SetName(name string) {
	p.Name = name
}

type MySlice []int

func (sl *MySlice) Add(val int) {
	*sl = append(*sl, val)
}

func (sl *MySlice) Count() int {
	return len(*sl)
}

func DemoMethods() {
	pers := Person{1, "Gosha"}
	pers.SetName("Gosha Golang")
	// (&pers).SetName("Gosha Golang")
	// fmt.Printf("updated person: %#v\n", pers)

	var acc Account = Account{
		Id:   1,
		Name: "Gosha",
		Person: Person{
			Id:   2,
			Name: "Gosha Golang",
		},
	}

	acc.SetName("golang.gosha")
	acc.Person.SetName("Test")

	// fmt.Printf("%#v\n", acc)

	sl := MySlice([]int{1, 2})
	sl.Add(5)
	fmt.Println(sl.Count(), sl)
}
