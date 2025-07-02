package loop

import "fmt"

func main() {
	// цикл без условия, while(true) OR for(;;;)
	for {
		fmt.Println("loop iteration")
		break
	}

	// цикл без условия, while(isRun)
	isRun := true
	for isRun {
		fmt.Println("loop iteration with condition")
		isRun = false
	}
}