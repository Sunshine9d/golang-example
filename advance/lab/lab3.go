package lab

import (
	"fmt"
	"sync"
	"time"
)

func runLab3() {
	var wg sync.WaitGroup
	//todo
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		fmt.Println("Inside go routine")
		//todo
		wg.Done()
	}()
	wg.Wait()
}
