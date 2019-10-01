package main

import (
	"fmt"
	"os"
)

func assertion(i interface{}) {
	s, ok := i.(int)
	fmt.Println(s, ok)
}

type error interface {
	Error() string
}
type PathError struct {
	Err  error
	Op   string
	Path string
}

func (p *PathError) Error() string {
	return p.Op + " " + p.Path + " " + p.Err.Error()
}

func mainError() {
	var j interface{} = 34
	assertion(j)
	f, err := os.Open("src/ChannelAndGoRoutine.go")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at Path " + err.Path + " failed to Open")
		return
	}
	fmt.Println(f.Name() + " Opened Successfully")
}
