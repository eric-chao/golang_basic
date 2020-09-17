package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Printf("GOMAXPROCS is %d\n", runtime.GOMAXPROCS(0))

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		num := i
		go func() {
			a := 0
			for i := 0; i < 1e8; i++ {
				a += 1
			}
			fmt.Printf(" result: %d -- %d \n", num, a)
			wg.Done()
		}()
	}
	wg.Wait()
}

func init() {
	runtime.GOMAXPROCS(2)
}