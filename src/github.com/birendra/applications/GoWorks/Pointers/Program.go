package main

import (
	"fmt"
)

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
func fun() *int {
	v := 1
	return &v
}
func fun1(add *int) int {
	*add++
	return *add
}
func main() {
	v := 34
	p := &v
	fmt.Printf("%v\n", *p)
	*p = 234
	fmt.Printf("%v %p\n", *p, p)
	fmt.Printf("%p\n", fun())
	fmt.Printf("%v\n", fun1(&v))
	fmt.Printf("%v\n", v)
	// new Function
	p1 := new(float64)
	fmt.Printf("%v %p\n", *p1, p1)
	*p1 = 123324.354435534
	fmt.Printf("%v %p\n", *p1, p1)
	fmt.Printf("%v\n", gcd(1000, 2500))
	fmt.Printf("%v\n", fib(4000))

}
