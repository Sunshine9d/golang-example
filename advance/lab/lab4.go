package lab

import (
	"fmt"
	"sync"
)

func runLab4() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		ch <- 1
		wg.Done()
		//fixed
		close(ch)
	}()
	wg.Wait()
	val := <-ch
	fmt.Println(val)
	//To fix the deadlock issue, you need to ensure that the channel is
	//closed after all goroutines have finished their work. You can move the close(ch) statement inside the goroutine. Here is the corrected code:
	close(ch)
	//<-ch
}
