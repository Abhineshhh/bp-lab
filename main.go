package main

import (
	"bp-lab/internal"
	"fmt"
)

func main() {
	fmt.Printf("\nHii ! Lets Implement backpressure in Streaming APIs \n")

	internal.BlockingPipeline() // Block Strategy
	internal.DroppingPipeline() // Drop Strategy
	internal.BufferedPipeline() // Buffer Strategy
	internal.BatchPipeline()    // Reshape Strategy - Batch

}
