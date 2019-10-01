package main

import (
	"TutorialLedge/tutorialedge"
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Main")
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go tutorialedge.MyGoroutine(&waitGroup)
	waitGroup.Wait()
	fmt.Println("Finished")
	waitGroup.Add(1)
	go func() {
		fmt.Println("I am Anonymous")
		waitGroup.Done()
	}()
	waitGroup.Wait()

	//Api Access
	tutorialedge.HandleRequest()
}
