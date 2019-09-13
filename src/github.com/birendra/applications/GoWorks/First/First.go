package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var i, j int = 5, 6
	c := i + j
	fmt.Println(c)
	var k string = "hello"
	var str string = strconv.Itoa(c)
	fmt.Printf("%v %T\n", str, str)
	fmt.Printf("%s\n", k)
	var cal float64 = float64(i) * float64(j) / float64(c)
	m, n := float64(i), float64(j)
	var min float64 = math.Min(float64(i), float64(j))
	fmt.Printf("%v %v %v %v\n", cal, min, n, m)
	a := 10
	b := 3
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a & ^b)
	//String
	var str1 string = "Hello!, World"
	str2 := strconv.Itoa(a)
	arrayByte := []byte(str1)
	fmt.Printf("%v ,%T,%v\n", arrayByte, arrayByte, str2)
	//rune
	var xx rune = 'a' //int32 not utf-8 like string
	fmt.Printf("%v, %T\n", xx, xx)
}
