package main

import (
	"demo/pkg2"
	"fmt"
)

func main() {

	//declaring and using implicit datatype
	//Basic syntax with declaring any variable => var variableName dataType. var is a keyword which is used to declare one or more variables
	/*var iVal int = 10
	iVal2 := 20
	fmt.Println(iVal, iVal2)

	//declaring and using array
	var arr [3]int = [3]int{10, 25, 35}
	fmt.Println(arr)

	//declaration with initialisation
	var arr3 [10]int
	arr3[1] = 10
	fmt.Println(arr3)

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
	fmt.Println(myVar)*/

	//working with loops
	//var myArr [5]int = [5]int{10, 20, 30, 40, 50}

	//infinite loop
	/*for {
		myArr[0] = 100
		fmt.Println("DO NOT RUN THIS CODE AS THIS IS A INFINITE LOOP!!!!!!!")
	}*/

	iVal := 10

	//this is like a while loop
	for iVal <= 20 {

		iVal += 5
	}

	fmt.Println("iVal ", iVal)

	for i := 0; i < 10; i++ {
		iVal += 1
	}

	fmt.Println("iVal ", iVal)

	//working with loops on collections
	var myArr [5]int = [5]int{10, 20, 30, 40, 50}

	for index, value := range myArr {
		fmt.Println("Value from my array at index ", index, " is ", value)
	}

	pkg2.MyPrint()
	//fmt.Println("This array is also from package 2", pkg2.Pkg2Arr)

	var customStructVar pkg2.MyStruct
	customStructVar.Name = "abc"

	customStructVar.PrintMyStruct()

}

//defining functions in go
//func functionName (parameters) (return values) {
//	function body
//}
