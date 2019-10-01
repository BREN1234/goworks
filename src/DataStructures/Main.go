package main

import (
	"DataStructures/datastructure"
	"fmt"
)

func main() {
	items := []int{12, 3, 2, 23, 53, 546, 6, 75, 76, 56, 544, 53, 534, 232}
	fmt.Println(datastructure.LinearSearch(items, 23))
	//selection Sort
	sortedSlice := datastructure.SelectionSort(items)
	fmt.Println(sortedSlice)
	listOfItems := datastructure.RandomNumberGenerator(20)
	fmt.Printf("Random Number: %v\n", listOfItems)
	fmt.Printf("After Selection Sort: %v\n", datastructure.SelectionSort(listOfItems))
}
