package channelworks

func GetDigit(n int, ch chan int) {
	for n != 0 {
		digit := n % 10
		ch <- digit
		n = n / 10
	}
	close(ch)
}
func Sq1(n int, ch chan int) {
	sum := 0
	chh := make(chan int)
	go GetDigit(n, chh)
	for digits := range chh {
		sum = sum + digits*digits
	}
	ch <- sum
}
func Cub(n int, ch chan int) {
	sum := 0
	chh := make(chan int)
	go GetDigit(n, chh)
	for digits := range chh {
		sum = sum + digits*digits*digits
	}
	ch <- sum
}
func Square(n int, ch chan int) {
	sum := 0
	for n != 0 {
		digits := n % 10
		n = n / 10
		sum = sum + digits*digits
	}
	ch <- sum

}
func Cube(n int, ch chan int) {
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
func Producer(ch chan int) {
	for i := 0; i < 20; i++ {
		ch <- i
	}
	close(ch)
}
