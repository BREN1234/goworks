package learning

import (
	"fmt"
	"math"
	"time"
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

type Circle struct {
	R float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.R, 2)
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}
func (s Square) Perimeter() float64 {
	return 4 * s.Side
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
type Object interface {
	Volume() float64
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
func (r Cuboid) Volume() float64 {
	return r.Base.X * r.Base.Y * r.H
}

//Empty Interface(dynamic type and  Value)
func Explain(i interface{}) {
	fmt.Printf("the value given to function is type of '%T' with value %v\n", i, i)
}

type MyString string

//From Video on Interface part 2
type Notifier interface {
	Notify()
}
type User struct {
	Name  string
	Email string
}

func (u *User) Notify() {
	fmt.Printf("Sending User Email Notification to %s <%s>\n", u.Name, u.Email)
}
func SendNotification(n Notifier) {
	n.Notify()
}

//Part3 interface
type Printer interface {
	Print()
}

//implementing the Printer interface in User
func (u User) Print() {
	fmt.Printf("UserName: %s and Email : %s\n", u.Name, u.Email)
}
func CalculateTotalArea(s ...Shape) float64 {
	sum := 0.0
	for _, T := range s {
		time.Sleep(1 * time.Second)
		fmt.Println("area", T.Area())
		sum = sum + T.Area()

	}
	return sum
}
