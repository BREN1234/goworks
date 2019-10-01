package main

import (
	"Interfaces/interfacework"
)

func main() {
	listOfObjects := []interfacework.Speaker{
		interfacework.Dog{
			Name:     "Dog",
			IsMammal: true,
		},
		interfacework.Bird{
			Name:     "Bird",
			IsMammal: false,
		},
		interfacework.Cat{
			Name:     "Cat",
			IsMammal: true,
		},
	}
	for _, value := range listOfObjects {
		value.Speak()
	}

	var objectCat interfacework.Speaker = &interfacework.Cat{Name: "cat", IsMammal: true}
	objectCat.Speak()
	var ob interfacework.Cat = interfacework.Cat{Name: "Caat", IsMammal: true}
	ob.Speak1()

}
