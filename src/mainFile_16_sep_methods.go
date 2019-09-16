package main

import (
	"fmt"
)

//Methods in Golang
//Prototype
//func (receiver_name Type) method_name(parameter_list) (return Type){
//	Code
//}
//Employee
type employee struct {
	name       string
	age        int
	height     float64
	department string
	salary     float64
}

func (e employee) show() {
	fmt.Println(e.name)
	fmt.Println(e.age)
	fmt.Println(e.height)
	fmt.Println(e.salary)
	fmt.Println(e.department)

}

//Rectangle
type rectangel struct {
	x, y float64
}

func (ptr *rectangel) area() float64 {
	return ptr.x * ptr.y
}
func (rec rectangel) perimeter() float64 {
	return rec.x * rec.y
}

//Type Definition
type data int

//Defing methods
func (d data) mult(d2 data) data {
	return d * d2
}
func main() {
	employee1 := employee{
		name:       "Birendra",
		age:        26,
		height:     150.0,
		department: "Analytics",
		salary:     12.0,
	}
	employee1.show()
	rec := rectangel{
		x: 20.0,
		y: 10.0,
	}
	fmt.Println(rec.perimeter())
	recPtr := &rec
	fmt.Println(recPtr.area())
	value1 := data(23)
	value2 := data(10)
	fmt.Println(value1.mult(value2))
}
