package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CelsiusToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func FahrenheitToC(F Fahrenheit) Celsius {
	return Celsius((F - 32) * 5 / 9)
}
func main() {
	map1 := map[string]int{
		"a": 1, "b": 2, "c": 3,
	}
	fmt.Printf("%v\n", map1)
	fmt.Printf("%v\n", map1["a"])
	map2 := map[[3]string]int{}
	fmt.Printf("%v\n", map2)
	map3 := make(map[string]int)
	map3 = map[string]int{"v": 123, "g": 4353}
	fmt.Printf("%v\n", map3)
	map3["added"] = 123214
	fmt.Printf("%v\n", map3)
	//delete the map
	delete(map1, "a")
	fmt.Printf("%v\n", map1)
	//Checking the keys is in map
	key, ok := map1["a"]
	fmt.Printf("%v %v len %v\n", key, ok, len(map1))
	//Maps are pass by reference
	map1_1 := map1
	delete(map1_1, "c")
	fmt.Printf("%v %v\n", map1_1, map1)
	//Type Declearation
	fmt.Printf("%v\n", FahrenheitToC(1234.343))

}
