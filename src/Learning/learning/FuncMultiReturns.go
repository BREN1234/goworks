package learning

import "fmt"

func init() {
	fmt.Println("Mult-Return Function package is initilaze")
}
func MultReturns(l float64, b float64) (float64, float64) {
	return l * b, 2 * (l + b)
}
func NamedValueReturn(l float64, b float64) (area, perimeter float64) {
	area = l * b
	perimeter = 2 * (l + b)
	return
}
