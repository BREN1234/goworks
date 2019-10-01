package main

import "fmt"

func getDigit(n int, ch chan int) {
	for n != 0 {
		digit := n % 10
		ch <- digit
		n = n / 10
	}
	close(ch)
}
func sq1(n int, ch chan int) {
	sum := 0
	chh := make(chan int)
	go getDigit(n, chh)
	for digits := range chh {
		sum = sum + digits*digits
	}
	ch <- sum
}
func cub(n int, ch chan int) {
	sum := 0
	chh := make(chan int)
	go getDigit(n, chh)
	for digits := range chh {
		sum = sum + digits*digits*digits
	}
	ch <- sum
}
func square(n int, ch chan int) {
	sum := 0
	for n != 0 {
		digits := n % 10
		n = n / 10
		sum = sum + digits*digits
	}
	ch <- sum

}
func cube(n int, ch chan int) {
	sum := 0
	for n != 0 {
		digits := n % 10
		n = n / 10
		sum = sum + digits*digits*digits
	}
	ch <- sum

}
func SendOnly(ch chan<- int) { //Send only cahnnel
	ch <- 1234

}
func producer(ch chan int) {
	for i := 0; i < 20; i++ {
		ch <- i
	}
	close(ch)
}
func mainChannel() {
	squareChannel := make(chan int)
	cubeChannel := make(chan int)
	var number int = 1234
	go square(number, squareChannel)
	go cube(number, cubeChannel)
	sq, cu := <-squareChannel, <-cubeChannel
	fmt.Println("Sum is ", sq+cu)

	ch := make(chan int)
	go SendOnly(ch)
	fmt.Println(<-ch)

	chh := make(chan int)
	go producer(chh)
	for {
		v, ok := <-chh
		if ok == false {
			break
		}
		fmt.Println(v)
	}
	//range loop
	chhh := make(chan int)
	go producer(chhh)
	for v := range chhh {
		fmt.Println(v)
	}

	chn := make(chan int)
	chn1 := make(chan int)
	go sq1(123, chn)
	go cub(123, chn1)
	sqValue := <-chn
	cubValue := <-chn1
	fmt.Println(sqValue, cubValue)
}
