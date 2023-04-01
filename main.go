package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter input")
	inputBuffer := bufio.NewReader(os.Stdin)
	input, _ := inputBuffer.ReadString('\n')
	inputUpper := strings.ToUpper(input)
	fmt.Println(inputUpper)
}
