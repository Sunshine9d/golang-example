package routine

import (
	"fmt"
	"sync"
	"time"
)

func start() {
	go process()
	fmt.Println("Starting")
}

func process() {
	fmt.Println("Processing")
}

func calculateSquare(x int) int {
	time.Sleep(1 * time.Second)
	square := x * x
	fmt.Println(square)
	return square
}

func Calculate() {
	start_time := time.Now()
	for i := 0; i < 2; i++ {
		go calculateSquare(i)
	}
	elapsed := time.Since(start_time)
	//Wait for 2 seconds to allow all the go routines to finish
	time.Sleep(2 * time.Second)
	fmt.Printf("Took %s\n", elapsed)
}

func calculateSquare1(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	square := i * i
	fmt.Println(square)
}

func runWaitGroup() {
	var wg sync.WaitGroup
	start_time := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go calculateSquare1(i, &wg)
	}
	elapsed := time.Since(start_time)
	wg.Wait()
	fmt.Printf("Took %s\n", elapsed)
}

func RunAdvance() {
	//Calculate()
	//start()
	//time.Sleep(2 * time.Second)
	runWaitGroup()
}
