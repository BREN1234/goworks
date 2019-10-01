package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
)

func Counter(ptr, ptr1 *[32]byte) int {
	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println(ptr1)
	fmt.Println(*ptr1)
	if len(*ptr) != len(*ptr1) {
		return 0
	}
	counter := 0
	for index, value := range *ptr {
		if value != ptr1[index] {
			return counter
		}
		counter++
	}
	return counter
}
func mainchap4() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	//Counting number bits different in sha256
	count := Counter(&c1, &c2)
	fmt.Println(count)
	//runes
	var runes []rune
	for i, r := range "Hello World" {
		runes = append(runes, r)
		fmt.Println(i, string(r))
	}
	fmt.Println(string(runes))
	fmt.Printf("%q\t", runes)
	var input = bufio.NewReader(os.Stdin)
	text, g := input.ReadString('\n')
	fmt.Println(text, g)

}
