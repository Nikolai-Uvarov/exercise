package worker

import (
	"fmt"
	"sync"
)

func Start(i int, c chan int, wg *sync.WaitGroup) {
	fmt.Println("Worker ", i, " started!")

	wg.Add(1) //портим код для неудачного коммтита и последующего отката
	//этот код перенесем внутрь горутины и получим нестабильную непредсказуемую работу программы

	c <- 1
	wg.Done()
}

func Counter(wg *sync.WaitGroup) (chan int, *int) {

	c := make(chan int)
	var i *int = new(int)
	*i = 0

	go func(c chan int, i *int) {
		for m := range c {
			*i += m
		}
		wg.Done()
	}(c, i)

	return c, i
}
