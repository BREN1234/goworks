package main

import (
	"fmt"
	"log"
)

var length, breadth float64 = 4, 5

func init() {
	println("main package initialized")
	if length <= 0 {
		log.Fatal("Length cannot be zero or negative")
	}
	if breadth <= 0 {
		log.Fatal("Breadth cannot be zero or negative")
	}
}
func sub(s []int) {
	for i := range s {
		s[i] -= 2
	}
}

//AnnonymousFunction
func sq() func() int {
	var x int
	fmt.Println("X= ", x)
	return func() int {
		fmt.Println("X= ", x)
		x++
		return x * x
	}
}
func fun1(i func(p, q string) string) {
	fmt.Println(i("Passing Anonymous function ", "as argument into another function"))
}
func main() {
	/**fmt.Println(learning.MultReturns(3, 4))
		fmt.Println(learning.MultReturns(0, 4))
		fmt.Println(learning.MultReturns(3, 0))
		fmt.Println(learning.MultReturns(3, -1))
		fmt.Println(learning.NamedValueReturn(3, 4))
		fmt.Println(learning.ReverseSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 0}))
	outer:
		for i := 0; i < 10; i = i + 1 {
			for j := 0; j < i; j++ {
				fmt.Println(i, j)
				if i == j {
					break outer
				}
			}
		}
		var aray []int = []int{1, 2, 3, 4, 5, 6, 7}
		for i := 0; i < len(aray); i++ {
			fmt.Println(aray[i])
		}
		var sum int = 0
		for _, v := range aray {
			sum += int(v)

		}
		fmt.Println(sum)*/
	darr := [...]int{1, 2, 3, 4, 5, 56, 6, 0, 99}
	dslice := darr[5:]
	fmt.Println(dslice)
	for i := range dslice {
		fmt.Printf("%v ", i)
		dslice[i]++
	}
	fmt.Println()
	fmt.Println(darr, cap(dslice))
	//creating the slice using make
	sl := make([]int, 5, 5)
	fmt.Println(sl)
	//Appending Slice
	sl = append(sl, 2)
	fmt.Println(sl, len(sl), cap(sl))
	sl = append(sl, 1, 2, 3, 4, 5, 6)
	fmt.Println(sl, len(sl), cap(sl))
	sl = append(sl, []int{11, 22, 33, 44}...)
	fmt.Println(sl, len(sl), cap(sl))

	var names []string
	if names == nil {
		names = append(names, []string{"ram", "syam"}...)
	}
	fmt.Println(names, len(names), cap(names))
	fmt.Println(dslice)
	sub(dslice)
	fmt.Println(dslice)
	//Anonymous Functions
	f := sq()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	func(ele string) {
		fmt.Println(ele)
	}("Hello! Passing Arguments to Anonymous Function")

	fu := func() {
		fmt.Println("Anonymous function assigned to a variable")
	}
	fu()

	//passing Anonymous function as arguments into another function
	fu1 := func(p, q string) string {
		return p + q + "Hel"
	}
	fun1(fu1)
	gh := []string{"Hello!", "world"}
	change(gh...)
	fmt.Println(gh)
	//maps
	var map1 map[string]int
	if map1 == nil {
		fmt.Println("creating Maps")
		map1 = make(map[string]int)
	}
	map1["Name"] = 1
	map1["age"] = 1
	fmt.Println(map1)
	delete(map1, "asdfasdf ")
	fmt.Println(map1)
	nj := intPrint("Hello Señor") //ñ requires more than one byte do rune handles
	fmt.Println(nj)
	for i := 0; i < len(nj); i++ {
		fmt.Printf("%c\t", nj[i])
		fmt.Printf("%x\t\n", nj[i])
		fmt.Printf("%v\n", nj[i])
	}
	nk := intPrintRune("Hello Señor")
	fmt.Println(nk)
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9} //slice of bytes
	str := string(byteSlice)
	fmt.Println(str)

}
func change(s ...string) {
	s[0] = "Go"
	s = append(s, "asdfsd")
	fmt.Println(s)
}
func intPrint(s string) []byte {
	var sl []byte = make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		sl[i] = s[i]
	}
	return sl
}
func intPrintRune(s string) []rune {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		r[i] = rune(s[i])
	}
	return r
}
