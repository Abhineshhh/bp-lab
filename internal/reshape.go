package internal

import (
	"fmt"
	"time"
)

// Don't react to pressure - prevent it. Group items into batches before sending. Producer pace is controlled.
func BatchPipeline() {

	fmt.Printf("\nImplementing Batching strategy : \n")

	in := make(chan int, 20)
	out := make(chan []int, 5)

	// producer - sends individual items fast
	go func() {
		defer close(in)

		for i := 1; i <= 12; i++ {
			in <- i
		}
	}()

	// batcher - groups into chunks of 4
	go func() {
		defer close(out)

		var batch = make([]int, 0, 4)

		for v := range in {
			batch = append(batch, v)

			if len(batch) == 4 {
				out <- batch
				batch = make([]int, 0, 4)
			}
		}

		// flush remaining
		if len(batch) > 0 {
			out <- batch
		}
	}()

	// consumer - recieve batches
	for b := range out {
		fmt.Println("batch : ", b)
		time.Sleep(200 * time.Millisecond)
	}
}
