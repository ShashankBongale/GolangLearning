package main

import (
	"fmt"
)

func main() {

	//declaring and using implicit datatype
	//Basic syntax with declaring any variable => var variableName dataType. var is a keyword which is used to declare one or more variables
	var iVal int = 10
	iVal2 := 20
	fmt.Println(iVal, iVal2)

	//declaring and using array
	var arr [3]int = [3]int{10, 25, 35}
	fmt.Println(arr)

	//declaring and using slice
	var mySlice []int = []int{-100, -200}
	mySlice = append(mySlice, 100, 200)
	fmt.Println(mySlice)

	//shorthand notation
	mySlice2 := []string{"abc", "xyz"}
	fmt.Println(mySlice2)
}
