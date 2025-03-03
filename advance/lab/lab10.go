package lab

import (
	"fmt"
	"sync"
)

func runLab10() {
	var wg sync.WaitGroup
	wg.Add(3)
	jobs := make(chan int)
	results := make(chan int)

	go func() {
		for j := range jobs {
			fmt.Println("received job", j)
			result := j * 2
			results <- result
		}
		// your code goes here
	}()

	go func() {
		for j := 1; j <= 3; j++ {
			jobs <- j
			fmt.Println("sent job", j)
		}
		// your code goes here
	}()

	go func() {
		for r := range results {
			fmt.Println("received result", r)
		}
		// your code goes here
	}()

	wg.Wait()
}
