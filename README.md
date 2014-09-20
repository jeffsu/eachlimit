async
=====

A simple implementation of node's popular [async#eachLimit](https://github.com/caolan/async#eachLimit) in Golang 


Install
=======

    go get github.com/jeffsu/eachlimit

Example
=======

    package main
    
    import (
    	"fmt"
    	"time"
    
    	"github.com/jeffsu/eachlimit"
    )
    
    func main() {
      limit := 3
    	el := eachlimit.New(limit)
    
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
