package pkg2

import "fmt"

func MyPrint() {

	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println("Recovering from panic, error: ", msg)
		}
	}()

	fmt.Println("This is my print from package2")

	dividend := 0
	iVal := 10 / dividend

	fmt.Println("Result of division ", iVal)
}

var Pkg2Arr [3]int = [3]int{10, 20, 30}
