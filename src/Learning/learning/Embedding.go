package learning

import "fmt"

type Admin struct {
	Person User
	Level  int64
}
type Show interface {
	Display()
}

//Embedding
type AdminEmbedded struct {
	User
	Level int64
}

func (a Admin) Display() {
	fmt.Printf("Admins Name: %s Admin Email: %s and Level: %d\n", a.Person.Name, a.Person.Email, a.Level)
}
func (a AdminEmbedded) Display() {
	fmt.Printf("Admins Name: %s Admin Email: %s and Level: %d\n", a.Name, a.Email, a.Level)
}
