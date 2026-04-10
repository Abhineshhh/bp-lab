package main

import (
	"bp-lab/internal"
	"fmt"
)

func main() {
	fmt.Println("Hii ! Lets Implement backpressure in Streaming APIs")

	// internal.BlockingPipeline() // Block Strategy
	internal.DroppingPipeline() // Drop Strategy
}
