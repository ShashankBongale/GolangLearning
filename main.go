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

	//declaring and using maps
	var myMap map[string]int = make(map[string]int)
	myMap["shashank"] = 1000
	fmt.Println(myMap)

	//map with slice as value
	var mapWithSlice map[string][]int = make(map[string][]int)
	mapWithSlice["shashank"] = append(mapWithSlice["shashank"], 300, 400)
	fmt.Println(mapWithSlice)

	for i := 0; i < 3; i++ {
		fmt.Println(arr[i])
	}

	//using struct
	var myStruct struct {
		name string
		id   int
	}

	myStruct.name = "shashank"
	myStruct.id = 100

	fmt.Println(myStruct)

	//declaring a user defined struct type
	type myType struct {
		name string
		sVal string
	}

	var myVar myType
	myVar.name = "Random"
	fmt.Println(myVar)

}
