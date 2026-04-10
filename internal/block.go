package internal

import (
	"fmt"
	"sync"
	"time"
)

// internals

// type hchan struct {
// 	qcount   uint           // items currently sitting in buf
// 	dataqsiz uint           // capacity - 0 for unbuffered
// 	buf      unsafe.Pointer // ring buffer (nil when dataqsiz == 0)
// 	sendx    uint           // next write index
// 	recvx    uint           // next read index
// 	recvq    waitq          // goroutines parked on recieve
// 	sendq    waitq          // goroutines parked on send <- backpressure queue
// 	lock     mutex
// }

// Producer waits until consumer is ready. Nothing is lost. Throughput = slowest side.

func BlockingPipeline() {

	fmt.Printf("\nImplementing Blocking strategy : \n")

	ch := make(chan int) // dataqsiz=0, buf=nil

	var wg sync.WaitGroup
	wg.Add(1)

	// Producer goroutine
	go func() {
		defer close(ch) // signals "no more data" to consumer
		for i := 0; i < 10; i++ {
			fmt.Printf("[producer] sending %d - will block until received\n", i)
			ch <- i
			//   runtime calls chansend1()
			//   if recvq is empty -> goroutine is parked in sendq
			//   when consumer calls chanrecv(), it dequeues us directly
			fmt.Printf("[producer] %d was taken - unblocked\n", i)
		}
	}()

	// Consumer goroutine
	go func() {
		defer wg.Done()
		for v := range ch { // range closes when ch is closed+drained
			fmt.Printf("[consumer] processing %d\n", v)
			time.Sleep(200 * time.Millisecond) // slow consumer
		}
	}()

	wg.Wait()
}

// The rendezvous happens inside chansend / chanrecv in the runtime. When both a sender and receiver arrive, the runtime copies the value directly from the sender's stack frame to the receiver's - no intermediate buffer at all. This is called a direct send.
