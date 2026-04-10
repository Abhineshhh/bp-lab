# Go Backpressure Strategies

This project demonstrates various backpressure handling techniques in Go using channels. Backpressure is a mechanism to cope with a fast producer and a slow consumer, preventing the system from being overwhelmed.

## Strategies Implemented

The following strategies are implemented as separate examples:

1.  **Blocking:** The producer blocks until the consumer is ready to receive the data. This is the default behavior for unbuffered channels. It ensures no data is lost, but the overall throughput is limited by the slowest component.

2.  **Dropping:** If the producer tries to send data to a full channel, the data is dropped instead of making the producer wait. This is a lossy but fast strategy, useful when processing the most recent data is more important than processing all data.

3.  **Buffering:** A channel with a buffer allows the producer to send a certain number of items before it blocks. This helps to absorb short bursts of data from the producer and smooth out the processing rate.

4.  **Batching (Reshaping):** Instead of sending individual items, the producer groups them into batches. This can be more efficient as it reduces the overhead of channel communication and allows the consumer to process items in chunks.

## How to Run

To run the examples, execute the following command from the root of the project:

```sh
go run main.go
```

This will run all the implemented pipeline strategies sequentially and print the output for each.
