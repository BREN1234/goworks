package learning

import (
	"fmt"
	"math"
)

type Employee struct {
	Name   string
	Salary float64
}

func (e Employee) Show() {
	fmt.Println(e.Name)
	fmt.Println(e.Salary)
}

//REctangel
type Rectangle struct {
	X, Y float64
}

func (r Rectangle) Area() float64 {
	return r.X * r.Y
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.X * r.Y)
}

//Pass by address or semantic reference
/*func (ptr *Rectangle) Area() float64 {
	return ptr.X * ptr.Y
}
func (ptr *Rectangle) Perimeter() float64 {
	return 2 * (ptr.X * ptr.Y)
}
*/
// built in types
type Data int

func (d Data) Mult(d1 Data) Data {
	return d * d1
}

type Point struct {
	X, Y float64
}

//Traditional function
func Dist(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//methods
func (p Point) Dist(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//Path
type Path []Point

func (path Path) Line() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum = sum + path[i-1].Dist(path[i])
		}
	}
	return sum
}

//Pointer methdos
func (ptr *Point) ScaleBy(scale float64) {
	ptr.X = ptr.X * scale
	ptr.Y = ptr.Y * scale
}

//Interface
type Shape interface {
	Area() float64
	Perimeter() float64
}
type Cuboid struct {
	Base Rectangle
	H    float64
}

func (r Cuboid) Area() float64 {
	return r.Base.Area() * r.H
}

func (r Cuboid) Perimeter() float64 {
	return 2 * ((r.Base.X * r.Base.Y) + (r.Base.Y * r.H) + (r.H * r.Base.X))
}