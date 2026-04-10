package internal

import (
	"fmt"
	"time"
)

// Producer runs freely up to cap N. Then blocks (not drops). Good for absorbing short bursts.
func BufferedPipeline() {

	fmt.Printf("\nImplementing Buffering strategy : \n")

	ch := make(chan int, 4) // buffer of 4

	// producer
	go func() {

		defer close(ch)

		for i := 1; i <= 10; i++ {
			ch <- i // free for first 4, then BLOCKS
			fmt.Printf("sent %d (queue : %d/4)\n", i, len(ch))

		}

	}()

	// consumer
	for v := range ch {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("got ", v)
	}
}
