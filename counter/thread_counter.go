package counter

import (
	"fmt"
	"io"
	"sync"
)

var waitGroup sync.WaitGroup
var completedMsg = "Done."

// NonThreadedCount performs a group of counts using a single thread
func NonThreadedCount(countTarget int, writer io.Writer) {
	for k := 0; k < 4; k++ {
		for i := 0; i < countTarget; i++ {
			// do nothing
		}
		printMessage("Completed one non-threaded count", writer)
	}
	printMessage(completedMsg, writer)
}

// ThreadedCount performs a group of counts using goroutines
func ThreadedCount(countTarget int, writer io.Writer) {
	for i := 0; i < 4; i++ {
		waitGroup.Add(1)
		go increment(countTarget)
	}
	waitGroup.Wait()
	printMessage(completedMsg, writer)
}

// incrementally counts up to a given number
func increment(targetCount int) {
	defer waitGroup.Done()
	for i := 0; i < targetCount; i++ {
		// do nothing
	}
	fmt.Println("Completed one threaded count")
}

func printMessage(msg string, writer io.Writer) {
	_, _ = fmt.Fprintln(writer, msg)
}
