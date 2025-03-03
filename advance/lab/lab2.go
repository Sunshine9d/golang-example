package lab

import (
	"fmt"
	"sync"
)

// What should be the minimum size of the channel
// in the program to prevent any deadlock error?
// RunLab is a function.
func runLab() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 10)
	go produce(ch, &wg)
	wg.Wait()
}

func produce(ch chan int, s *sync.WaitGroup) {
	defer s.Done()
	for i := 0; i < 100; i += 10 {
		ch <- i

	}
	fmt.Println("Exiting Produce")
	close(ch)
	go consume(ch, s)
}

func consume(ch chan int, s *sync.WaitGroup) {
	defer s.Done()
	for val := range ch {
		fmt.Println("Received", val)
	}
	fmt.Println("Exiting Consume")
}
