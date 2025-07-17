package counter

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

// NonThreadedCount performs a group of counts using a single thread
func NonThreadedCount(countTarget int) {
	for k := 0; k < 4; k++ {
		for i := 0; i < countTarget; i++ {
			// do nothing
		}
		fmt.Println("Completed one non-threaded count")
	}
}

// ThreadedCount performs a group of counts using goroutines
func ThreadedCount(countTarget int) {
	for i := 0; i < 4; i++ {
		waitGroup.Add(1)
		go increment(countTarget)
	}
	waitGroup.Wait()
}

// incrementally counts up to a given number
func increment(targetCount int) {
	defer waitGroup.Done()
	for i := 0; i < targetCount; i++ {
		// do nothing
	}
	fmt.Println("Completed one threaded count")
}
