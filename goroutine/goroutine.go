package goroutine

import (
	"fmt"
	"sync"
)

func ImplementingGoRoutine() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		fmt.Println("Hi this is a goroutine")
		wg.Done()
	}()

	wg.Wait()

}
