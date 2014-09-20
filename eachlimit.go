package eachlimit

import "sync"

type EachLimit struct {
	wg sync.WaitGroup

	working    chan int
	concurrent int
}

func New(concurrent int) *EachLimit {
	el := new(EachLimit)
	el.concurrent = concurrent
	el.working = make(chan int, concurrent)
	return el
}

func (el *EachLimit) Wait() (waited bool) {
	waited = false

	waited = len(el.working) >= el.concurrent
	el.wg.Add(1)
	el.working <- 1
	return waited
}

func (el *EachLimit) WaitAll() {
	el.wg.Wait()
}

func (el *EachLimit) Done() {
	<-el.working
	el.wg.Done()
}
