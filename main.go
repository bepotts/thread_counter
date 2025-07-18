package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"thread_counter/counter"
	"time"
)

func main() {
	fmt.Println("Program Starting....")
	fmt.Print("What would you like to count to: ")
	count := readLine()
	fmt.Printf("Counting to: %s four times....\n", count)
	intCount := stringToInt(count)

	fmt.Println("Beginning non-threaded count...")
	ticker := time.Now()
	counter.NonThreadedCount(intCount, os.Stdout)
	duration := time.Since(ticker)
	fmt.Printf("Non-threaded count took %s\n", duration)

	fmt.Println("Starting thread count")
	ticker = time.Now()
	counter.ThreadedCount(intCount, os.Stdout)
	duration = time.Since(ticker)

	fmt.Printf("Threaded count took %s\n", duration)
}

// Reads a line from user input and returns it
func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		log.Fatalln(err)
	}
	return input
}

func stringToInt(inputString string) int {
	num, err := strconv.Atoi(inputString)
	if err != nil {
		log.Fatalln(err)
	}
	return num
}
