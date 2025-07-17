package main

import (
	"fmt"
	"thread_counter/counter"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	fmt.Println("Program Starting....")
	fmt.Println("Beginning non-threaded count...")
	ticker := time.Now()
	counter.NonThreadedCount(1000000)
	duration := time.Since(ticker)
	fmt.Println("Done.")
	fmt.Printf("Non-threaded count took %s seconds", duration)
}
