package counter

import (
	"fmt"
)

var threadedCount = 0

func NonThreadedCount(countTarget int) {
	var nonThreadedCount = 0
	for i := 0; i < countTarget; i++ {
		nonThreadedCount++
	}
}

func ThreadedCount(countTarget int) {

	fmt.Println("Starting threaded count")

}
