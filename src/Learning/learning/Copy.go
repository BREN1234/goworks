package learning

import "fmt"

func main0() {
	a := []int{1, 2, 3}
	b := a //Reference
	c := []int{}
	d := make([]int, 3, 4)
	fmt.Printf("%v %v\n", a, b)
	copy(d, a)
	fmt.Println("Copied=", a, c, d, len(d), cap(d))
	map1 := map[string]bool{"a": true, "b": false}
	map2 := make(map[string]bool)
	fmt.Println(map1, map2)
	for k, v := range map1 {
		map2[k] = v
	}
	fmt.Println(map1, map2)

}
