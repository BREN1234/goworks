package main

import (
	"Defer/deferworks"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:])
	deferworks.Increment()
	deferworks.LIFO()
}
