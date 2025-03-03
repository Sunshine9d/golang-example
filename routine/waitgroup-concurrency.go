package routine

import (
	"fmt"
	"sync"
)

func runWaitGroupConcurrency() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go func() {
			//defer wg.Done()
			fmt.Println(i)
			wg.Done()
		}()
	}
	fmt.Println("Waiting for all goroutines to finish")
	wg.Wait()
}
