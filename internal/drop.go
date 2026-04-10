package internal

import (
	"fmt"
	"time"
)

// Producer never waits. If channel is full, item is thrown away. Fast but lossy.
func DroppingPipeline() {

	fmt.Printf("\nImplementing Dropping strategy : \n")

	ch := make(chan int, 3) // buffer of 3

	// producer
	go func() {
		defer close(ch)

		for i := 1; i <= 10; i++ {
			select {
			case ch <- i:
				fmt.Println("sent", i)
				time.Sleep(10 * time.Millisecond)
			default: // channel full -> drop
				fmt.Println("Dropped", i) // dropping policy can be determined -  (either newest or oldest)
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()

	// consumer
	for v := range ch {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("got", v)
	}
}
