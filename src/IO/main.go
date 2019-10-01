package main

import "IO/ioworks"

func main() {
	ioworks.CopyFile("main.go", "main1.go") //Bug
	ioworks.CopyFileDefer("main1.go", "main2.go")
}
