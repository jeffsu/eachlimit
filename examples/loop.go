package main

import (
	"fmt"
	"time"

	"github.com/jeffsu/eachlimit"
)

func main() {
	el := eachlimit.New(3)

	for i := 0; i < 10; i++ {
		el.Wait()
		go func(i int) {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
			el.Done()
		}(i)
	}

	el.WaitAll()
	fmt.Println("done")
}
