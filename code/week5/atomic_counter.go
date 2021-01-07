package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//the editor send a sync/atomic error in the importation, but this is working. So, if you have this error, please ignore this, because this error is for the editor configuration.

func main() {
	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 80; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("ops:", ops)
}
