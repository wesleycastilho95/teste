package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func novasgoroutines(i int) {
	wg.Add(i)

	for j := 0; j < i; j++ {

		x := j
		go func(i int) {
			fmt.Println("Eu sou a goroutine numero:", i)
			wg.Done()
		}(x)
	}

	wg.Done()
}

func main() {

	novasgoroutines(10)
	wg.Wait()

}
