package recover

import "fmt"



func DeferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend", err)
		}
	}()
	fmt.Println("Some userful work")
	panic("something bad happened")
	}

func main() {
	DeferTest()
}