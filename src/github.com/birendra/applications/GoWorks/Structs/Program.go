package main

import (
	"fmt"
)

type Doctor struct {
	name      string
	age       int
	education []string
}

func main() {
	a := Doctor{
		name: "Birendra",
		age:  25,
		education: []string{
			"VJA",
			"Roberts",
			"SSSIHL",
		},
	}
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n %v\n %v\n", a.name, a.education[0], a.age)
	stArray := make([]Doctor, 2)
	stArray1 := []Doctor{}

	stArray[0] = Doctor{
		name: "Birendra",
		age:  25,
		education: []string{
			"VJA",
			"Roberts",
			"SSSIHL",
		},
	}
	stArray[1] = Doctor{
		name: "Rai", age: 23, education: []string{"a", "b", "f"},
	}
	fmt.Printf("%v %v\n %v\n %v\n", len(stArray), cap(stArray), stArray[0], stArray[1])
	for i := 0; i < 2; i++ {
		n := Doctor{name: "aa" + string(i), age: 23 + 1, education: []string{"abd", "aada"}}
		stArray1 = append(stArray1, n)
	}
	fmt.Printf("%v %v\n", stArray1, cap(stArray1))
	//Implicit Struct
	aDoctor := struct{ name string }{name: "John"}
	//Pass by value
	aDoctorA := aDoctor
	aDoctorA.name = "Joy"
	fmt.Printf("%v\n %v\n", aDoctor, aDoctorA)
	//Pass by Reference
	aDoctorB := &aDoctor
	aDoctorB.name = "Changed"
	fmt.Printf("%v\n %v\n", aDoctor, aDoctorB)
}
