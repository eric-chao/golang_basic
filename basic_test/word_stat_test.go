package basic

import (
	"strings"
	"sync"
	"testing"
	"time"
)

func TestWordSplit(t *testing.T) {
	var s = "I am a good people I am a good person"

	var arr = strings.Split(s, " ")

	t.Logf("arr: %+v", arr)

	var i int32
	var ch = make(chan int32)
	var l sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(arr))

	var wordMap = make(map[string]int)
	for _, w := range arr {

		go func(s string) {
			l.Lock()
			defer func() {
				l.Unlock()
				wg.Done()
			}()
			if c, ok := wordMap[s]; ok {
				c = c + 1
				wordMap[s] = c
			} else {
				wordMap[s] = 1
			}
			i++
			ch <- i
		}(w)

	}

	go func() {
		for {
			select {
			case i, ok := <-ch:
				if !ok {
					t.Log("channel is closed.")
					return
				}
				t.Logf("task %d done ...", i)
			case <-time.After(time.Second):
				t.Log("task timeout ...")
				return
			//default:
			//	t.Log(" --- --- ")
			}
		}
	}()

	wg.Wait()
	close(ch)

	for k, v := range wordMap {
		t.Logf("word: %s -- count: %d", k, v)
	}
}
