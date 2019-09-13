package main

import (
	"fmt"
)

func main() {
	var array [3]int
	var arrayStr [2]string
	array1 := [3]int{1, 3, 4}
	array2 := [...]int{1, 2, 3, 4, 5, 6, 6, 5}
	fmt.Printf("%v\n", array)
	fmt.Printf("%v\n", array1)
	fmt.Printf("%v\n", array2)
	array[0] = 12
	arrayStr[0] = "Hello man"
	fmt.Printf("%v,%v,%v\n", arrayStr, array, len(array2))
	//Array Assignment are copy
	arrayCopy := array1
	arrayCopy[0] = 100
	fmt.Printf("Array1 = %v ArrayCopy= %v\n", array1, arrayCopy)
	//Pointer
	arrayCopy1 := &arrayCopy
	arrayCopy1[0] = 1000
	fmt.Printf(" ArrayCopy= %v ArrayCopy1 after addres assignment=%v\n", arrayCopy, arrayCopy1)

	//Matrix
	var matrix [3][3]int
	var matrix1 [3][3]int = [3][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}, [3]int{7, 8, 9}}
	fmt.Printf("%v\n %v %v\n", matrix, matrix1, len(matrix1))
	matrix[0][2] = 5
	matrix[1][1] = 5
	matrix[2][0] = 5
	fmt.Printf("%v\n", matrix)

	//Slice
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a := slice
	a[0] = 10 //Slice points to the same reference address insted of copy
	fmt.Printf("%v\n", a)
	b := slice[2:7]
	c := slice[4:6]
	fmt.Printf("%v, %v\n", b, c)
	fmt.Printf("%v %v \n", len(slice), cap(slice))
	slice1 := []int{}
	fmt.Printf("%v\n", slice1)
	fmt.Printf("cap= %v\n", cap(slice1))
	fmt.Printf("len= %v\n", len(slice1))
	slice1 = append(slice1, 1)
	fmt.Printf("%v\n", slice1)
	fmt.Printf("cap= %v\n", cap(slice1))
	fmt.Printf("len= %v\n", len(slice1))
	slice1 = append(slice1, 2, 3, 4, 5)
	fmt.Printf("%v\n", slice1)
	fmt.Printf("cap= %v\n", cap(slice1))
	fmt.Printf("len= %v\n", len(slice1))
	//case one               Decompose
	slice1 = append(slice1, []int{6, 7, 8}...)
	fmt.Printf("%v\n", slice1)
	fmt.Printf("cap= %v\n", cap(slice1))
	fmt.Printf("len= %v\n", len(slice1))
	//Common operation stack
	slice2 := []int{1, 2, 3, 4, 5, 6}
	//pop
	pop := slice2[1:]
	deleteLastElement := slice2[:len(slice2)-1]
	removingTheElementInBetween := append(slice2[:2], slice2[3:]...) // last
	fmt.Printf("%v %v %v\n", pop, deleteLastElement, removingTheElementInBetween)
	//using Make Func

	make1 := make([]int, 10, 100)
	fmt.Printf("%v\n", make1)
	fmt.Printf("%v %v\n", cap(make1), len(make1))

}
