package main

import (
	"io/ioutil"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/trace"
	"strconv"
	"sync"
)

func main() {
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

	runtime.GOMAXPROCS(4)

	var wg sync.WaitGroup
	for i := 0; i < 40; i++ {
		wg.Add(1)
		go func() {
			a := 0
			for i := 0; i < 1e6; i++ {
				a += 1
				if i == 1e6/2 {
					bytes, _ := ioutil.ReadFile(`README.MD`)
					inc, _ := strconv.Atoi(string(bytes))
					a += inc
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
