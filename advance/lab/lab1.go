package lab

import "fmt"

func consume1(ch chan int) {
	for {
		select {
		case num := <-ch:
			fmt.Printf("%d", num)
			break
		}
	}
}

func runLab1() {
	ch := make(chan int)
	go consume1(ch)
	for i := 0; i < 5; i++ {
		ch <- i
	}
}
