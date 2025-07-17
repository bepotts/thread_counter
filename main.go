package main

import (
	"fmt"
	"thread_counter/counter"
	"time"
)

func main() {
	fmt.Println("Program Starting....")
	fmt.Println("Beginning non-threaded count...")
	ticker := time.Now()
	counter.NonThreadedCount(1000000)
	duration := time.Since(ticker)
	fmt.Println("Done.")
	fmt.Printf("Non-threaded count took %s\n", duration)

	fmt.Println("Starting thread count")
	ticker = time.Now()
	counter.ThreadedCount(1000000)
	duration = time.Since(ticker)
	fmt.Println("Done.")
	fmt.Printf("Threaded count took %s\n", duration)
}
