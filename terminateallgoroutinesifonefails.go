package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	clients := 4
	// make it buffered, so all clients can fail without hanging
	notifyCh := make(chan struct{}, clients)
	go produce(50, ch, notifyCh)

	var wg sync.WaitGroup
	wg.Add(clients)
	for i := 0; i < clients; i++ {
		go func() {
			consumer(ch, notifyCh)
			wg.Done()
		}()
	}
	wg.Wait()

}

func consumer(in chan int, notifyCh chan struct{}) {
	fmt.Printf("Start consumer\n")
	for i := range in {
		if i == 20 {
			fmt.Printf("%d fails\n", i)
			notifyCh <- struct{}{}
			return
		} else {
			fmt.Printf("%d\n", i)
		}

	}
	fmt.Printf("Consumer stopped working\n")
}

func produce(N int, out chan int, notifyCh chan struct{}) {
	for i := 0; i < N; i++ {
		select {
		case out <- i:
		case <-notifyCh:
			close(out)
			return
		}
	}
	close(out)
}
