package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", BasicHandler)
	http.ListenAndServe("192.168.0.105:3000", nil)
}

func BasicHandler(writer http.ResponseWriter, reader *http.Request) {

	fmt.Fprintf(writer, "Hello")
}
