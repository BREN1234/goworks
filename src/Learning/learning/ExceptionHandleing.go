package learning

import "fmt"

func PanicAndRecover() {
	var action int
	fmt.Println("Enter 1 for Student and 2 for Professional")
	fmt.Scanln(&action)
	switch action {
	case 1:
		fmt.Println("You entered Student")
		break
	case 2:
		fmt.Println("You Entered Professional")
	default:
		panic(fmt.Sprintf("I am %d", action))
	}
	defer func() {
		action := recover()
		fmt.Println("Recover gave value", action)
	}()
}
