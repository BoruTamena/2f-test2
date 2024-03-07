package main

import (
	"fmt"
	"sync"
	"time"
)

func increment(num *int, by int, mu *sync.Mutex, ch chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < by; i++ {
		time.Sleep(time.Second)

		// Lock the mutex before incrementing the shared variable
		mu.Lock()

		*num = (*num) + 1

		// Unlock the mutex after modifying the shared variable
		mu.Unlock()

		if *num == 10 {
			ch <- true
			return
		}
	}
}

func main() {

	count := 5
	var wg sync.WaitGroup
	var mu sync.Mutex

	ch := make(chan bool)

	wg.Add(2)

	go increment(&count, 4, &mu, ch, &wg) // incrementing count by 4
	go increment(&count, 3, &mu, ch, &wg) // incrementing count by 3

	go func() {
		defer close(ch) // close the channel upon completion
		wg.Wait()
	}()

	full := <-ch
	fmt.Printf("full is %v\n", full)
	fmt.Println("Final count:", count)
}
