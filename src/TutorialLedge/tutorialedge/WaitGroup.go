package tutorialedge

import (
	"fmt"
	"sync"
)

func MyGoroutine(waitGroup *sync.WaitGroup) {
	fmt.Println("I am Go routine")
	waitGroup.Done()
}
