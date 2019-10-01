package interfacework

import (
	"fmt"
)

type Speaker interface {
	Speak()
}
type Dog struct {
	Name     string
	IsMammal bool
}
type Bird struct {
	Name     string
	IsMammal bool
}
type Cat struct {
	Name     string
	IsMammal bool
}

func (dog Dog) Speak() {
	fmt.Println(dog.Name, dog.IsMammal)
}
func (cat Cat) Speak() {
	fmt.Println(cat.Name, cat.IsMammal)
}
func (cat *Cat) Speak1() {
	fmt.Println(cat.Name, cat.IsMammal)
}
func (bird Bird) Speak() {
	fmt.Println(bird.Name, bird.IsMammal)
}
