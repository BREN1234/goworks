package main

import (
	"fmt"
	"time"
)

func myFunction(ch chan int) {
	fmt.Println(123 + <-ch)
}
func fun(strChannel chan string) {
	for i := 0; i < 10; i++ {
		strChannel <- "Hello World!"
	}
	close(strChannel)
}
func myBoolFunc(channel chan bool) {
	fmt.Println("I am from Booleam Channle")
	channel <- true
}
func mainRoutine() {
	fmt.Println("I am Main Goroutine")
	//Annoymous go functions
	/*go func() {
		fmt.Println("I am Annoymous Goroutine")
	}()
	*/
	time.Sleep(1 * time.Second)
	fmt.Println("Good Bye Main Function")
	myChannel1 := make(chan int)
	myBool := make(chan bool)
	go myBoolFunc(myBool)
	<-myBool
	go myFunction(myChannel1)
	myChannel1 <- 7

	strChannel := make(chan string)
	go fun(strChannel)
	for {
		res, ok := <-strChannel
		if ok == false {
			fmt.Println("Channel is closed")
			break
		}
		fmt.Println("Channel is open", res, ok)
	}
	strChannel1 := make(chan string)
	//Annoymous Funtion
	go func() {
		strChannel1 <- "Hi"
		strChannel1 <- "Hi"
		strChannel1 <- "Hi"
		strChannel1 <- "Hi"
		strChannel1 <- "Hi"
		strChannel1 <- "Hi"
		strChannel1 <- "Hi"
		strChannel1 <- "Hi"
		strChannel1 <- "Hi"
		strChannel1 <- "Hi"
		strChannel1 <- "Hi"
		close(strChannel1)
	}()
	for res := range strChannel1 {
		fmt.Println(res)
	}
}
