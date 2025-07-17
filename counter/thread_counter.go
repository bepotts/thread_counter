package counter

import "sync"

var waitGroup sync.WaitGroup

func NonThreadedCount(countTarget int) {
	for k := 0; k < 4; k++ {
		for i := 0; i < countTarget; i++ {
			// do nothing
		}
	}
}

func ThreadedCount(countTarget int) {
	for i := 0; i < 4; i++ {
		waitGroup.Add(1)
		go increment(countTarget)
	}
	waitGroup.Wait()
}

func increment(targetCount int) {
	defer waitGroup.Done()
	for i := 0; i < targetCount; i++ {
		// do nothing
	}
}
