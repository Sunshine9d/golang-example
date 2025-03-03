package lab

import "sync"

//Complete the code below to create two goroutines
//that will run concurrently using a WaitGroup.
//	The first goroutine should print the numbers 1 to 5,
//and the second goroutine should print the letters a to e.
//	The main function should wait for both goroutines to
//finish before exiting.

func runLab7() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers(&wg)
	go printLetters(&wg)
	wg.Wait()
}

func printLetters(wg *sync.WaitGroup) {
	for i := 'a'; i <= 'e'; i++ {
		println(string(i))
	}
	wg.Done()
}

func printNumbers(wg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		println(i)
	}
	wg.Done()
}
