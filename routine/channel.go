package routine

import (
	"fmt"
	"sync"
	"time"
)

func buy(ch chan string) {
	fmt.Println(`Waiting for data from the channel`)
	data := <-ch
	fmt.Println("Received data from the channel")
	fmt.Println(data)

}

func sell(ch chan string) {
	time.Sleep(5 * time.Second)
	ch <- "Furniture"
	fmt.Println("Sent data to the channel")
}

func run1() {
	ch := make(chan string)
	go buy(ch)
	go sell(ch)
	// Wait for 2 seconds to allow the go routines to finish
	time.Sleep(6 * time.Second)
}

func sell1(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- 1
	ch <- 2
	ch <- 3
	//ch <- 5
	go buy1(ch, wg)
	ch <- 4
	fmt.Println("Sent data to the channel")
	//wg.Done()
}

func buy1(ch chan int, wg *sync.WaitGroup) {
	//defer wg.Done()
	fmt.Println(`Waiting for data from the channel`)
	fmt.Println(<-ch)
	for i := range ch {
		fmt.Println(i)
		wg.Done()
	}
	fmt.Println("Received data from the channel")
}
func run2() {
	var wg sync.WaitGroup
	wg.Add(3)
	ch := make(chan int, 3)
	go sell1(ch, &wg)
	wg.Wait()
}

func RunChannel() {
	//run1()
	run2()
}
