package recover

import "fmt"



func DeferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend FIRST", err)
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend SECOND", err)
			// panic("second panic")
		}
	}()
	fmt.Println("Some userful work")
	panic("something bad happened")
}

