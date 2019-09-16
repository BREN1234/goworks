package learning

import "fmt"

func ReverseSlice(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
func main1() {

	s := []int{1, 2, 3, 4, 5}
	f := ReverseSlice(s)
	fmt.Printf("%v\n", f)
}
