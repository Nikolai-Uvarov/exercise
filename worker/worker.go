package worker

import "sync"

func start(c chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	c <- 1
	wg.Done()
}
