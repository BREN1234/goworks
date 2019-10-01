package deferworks

import (
	"fmt"
)

func Increment() {
	var i int = 0
	defer fmt.Println(i)
	i++
	return
}
func LIFO() {
	for i := 0; i <= 10; i++ {
		defer fmt.Println(i)
	}
}
