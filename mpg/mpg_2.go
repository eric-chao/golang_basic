package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

func main() {
	fmt.Printf("GOMAXPROCS is %d\n", runtime.GOMAXPROCS(0))

	fmt.Println("Hello")
}

func init() {
	runtime.GOMAXPROCS(2)

	// trace
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
}