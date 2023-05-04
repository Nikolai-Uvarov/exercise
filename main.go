package main

import (
	"fmt"
	"git_goroutines/worker"
	"sync"
)

var n int = 10

func main() {
	var wg sync.WaitGroup

	c, counter := worker.Counter(&wg)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker.Start(i, c, &wg)
	}

	wg.Wait()

	wg.Add(1)
	close(c)
	wg.Wait()

	fmt.Println("Work of ", n, "goroutines is finished!")
	fmt.Println("Counter is", *counter)
}
