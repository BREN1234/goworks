package main

import (
	"Channel/channelworks"
	"fmt"
)

func main() {
	squareChannel := make(chan int)
	cubeChannel := make(chan int)
	var number int = 1234
	go channelworks.Square(number, squareChannel)
	go channelworks.Cube(number, cubeChannel)
	sq, cu := <-squareChannel, <-cubeChannel
	fmt.Println("Sum is ", sq+cu)
	ch := make(chan int)
	go channelworks.SendOnly(ch)
	fmt.Println(<-ch)

	chh := make(chan int)
	go channelworks.Producer(chh)

	for {
		v, ok := <-chh
		if ok == false {
			break
		}
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	//range loop
	chhh := make(chan int)
	go channelworks.Producer(chhh)
	for v := range chhh {
		fmt.Printf("%v ", v)
	}
	fmt.Println()

	chn := make(chan int)
	chn1 := make(chan int)
	go channelworks.Sq1(123, chn)
	go channelworks.Cub(123, chn1)
	sqValue := <-chn
	cubValue := <-chn1
	fmt.Println(sqValue, cubValue)

	message := make(chan string)
	//signals := make(chan bool)
	select {
	case msg := <-message:
		fmt.Println("received Message:%s ", msg)
	default:
		fmt.Println("Not Received")
	}
	msg := "Message"
	select {
	case message <- msg:
		fmt.Println(msg)
	default:
		fmt.Println("No Message received")
	}
	fmt.Println()
}
