package main

import (
	"Learning/learning"
	"fmt"
)

func main() {
	var employee1 = learning.Employee{
		"Birendra",
		12.0,
	}
	employee1.Show()
	var rec = learning.Rectangle{
		20.0,
		10.0,
	}
	fmt.Println(rec.Perimeter())
	recPtr := &rec
	fmt.Println(recPtr.Area())
	//Non struct Type
	value1 := learning.Data(23)
	value2 := learning.Data(10)
	fmt.Println(value1.Mult(value2))
	//Points
	var p = learning.Point{4, 5}
	q := learning.Point{9, 5}
	fmt.Println(learning.Dist(p, q))
	fmt.Println(p.Dist(q))
	x := &learning.Point{3, 4}
	//scale using pointer
	x.ScaleBy(10)
	fmt.Println(*x)
	(*x).ScaleBy(10)
	fmt.Println(*x)
	(&p).ScaleBy(4)
	fmt.Println(p)
	p.ScaleBy(3) //implicit (&p)
	fmt.Println(p)
	//Path
	path := learning.Path{
		{1, 2}, {3, 4}, {5, 6}, {6, 7}, {7, 8},
	}
	fmt.Println(path.Line())
	//Interface
	var interf learning.Shape = learning.Cuboid{learning.Rectangle{2, 3}, 4}
	fmt.Println(interf)
	fmt.Println(interf.Area())
	fmt.Println(interf.Perimeter())
	interf = learning.Rectangle{3, 4}
	fmt.Println(interf)
	fmt.Println(interf.Area())
	fmt.Println(interf.Perimeter())
}
