package lab

import (
	"fmt"
	"sync"
)

func runLab9() {
	var wg sync.WaitGroup
	wg.Add(1)
	jobs := make(chan int, 3)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				wg.Done()
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	// your code goes here

	fmt.Println("sent all jobs")
	close(jobs)
	<-done
	wg.Wait()
}
