package main

import (
	"fmt"
	"sync"
	"time"
)

func increment(num *int, by int, ch chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	for i := 0; i < by; i++ {

		if *num == 10 {
			ch <- true
			return
		}
		*num = (*num) + 1
	}

}

func main() {

	count := 5
	var wg sync.WaitGroup

	ch := make(chan bool)

	wg.Add(2)

	go increment(&count, 4, ch, &wg)
	go increment(&count, 3, ch, &wg)

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	full := <-ch

	fmt.Printf("full is %v", full)
	fmt.Println(count)
}
