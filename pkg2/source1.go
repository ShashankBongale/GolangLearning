package pkg2

import "fmt"

func MyPrint() {
	defer fmt.Println("This is a defered statement")
	fmt.Println("This is my print from package2")
	//panic("Something bad happened")
}

var Pkg2Arr [3]int = [3]int{10, 20, 30}
