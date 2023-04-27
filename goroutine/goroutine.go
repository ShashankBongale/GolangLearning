package goroutine

import (
	"fmt"
	"sync"
)

func ImplementingGoRoutine() {

	var ch chan string = make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		fmt.Println("Sending message")
		ch <- "Task1"
	}()

	go func() {
		fmt.Println("Waiting for message")
		var str string = <-ch
		fmt.Println("Received message ", str)
		wg.Done()
	}()

	wg.Wait()

}
