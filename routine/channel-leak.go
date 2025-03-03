package routine

import "sync"

func runChannelLeak() {
	var wg sync.WaitGroup
	wg.Add(2)
	go leak(&wg)
	wg.Wait()
}

// The error in your code is due to the channel ch in the leak function not being closed,
// which causes the goroutine to wait indefinitely on val := <-ch.
//
//	This results in the WaitGroup never being marked as done,
//
// causing the wg.Wait() call in runChannelLeak to block forever.
//
//	To fix this, you need to close the channel ch after the goroutine reads from it. Here is the corrected code:
func leak(s *sync.WaitGroup) {
	ch := make(chan int)
	go func() {
		defer s.Done()
		val := <-ch
		println(val)
	}()
	defer s.Done()
}
