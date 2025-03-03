package routine

import (
	"fmt"
	"time"
)

func runTimeOut() {
	ch1 := make(chan int)
	go sendValue(ch1)
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timed out")
	}
}

func sendValue(ch1 chan int) {
	time.Sleep(3 * time.Second)
	ch1 <- 10
}
