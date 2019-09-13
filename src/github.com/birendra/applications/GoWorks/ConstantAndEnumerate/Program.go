package main

import (
	"fmt"
)

const (
	enumerate = iota << 3
	enumerate1
	enumerate2
)

func main() {
	const a = 42
	fmt.Printf("%v,%T\n", a, a)
	var b int32 = 8
	fmt.Printf("%v, %T\n", a+b, a+b)
	fmt.Printf("%v,%v,%v", enumerate, enumerate1, enumerate2)
}
