package routine

import (
	"fmt"
	"time"
)

func goOne(ch chan string) {
	ch <- "Hello"
}

func goTwo(ch chan string) {
	ch <- "World"
}

func runChannelSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go goTwo(ch2)
	go goOne(ch1)
	select {
	case data := <-ch1:
		fmt.Println("Received data from channel 1")
		fmt.Println(data)
		break
	case data := <-ch2:
		fmt.Println("Received data from channel 2")
		fmt.Println(data)
		break
		//default:
		//	fmt.Println("No data received")
	}
	time.Sleep(2 * time.Second)
}
