package main

import (
	"bufio"
	"fmt"
	"os"
)

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {
	//Reading Rune or UTF-8 Encoded
	reader := bufio.NewReader(os.Stdin)
	char, _, errr := reader.ReadRune()
	//Reading Full text
	readerFullText, _ := reader.ReadString('\n')
	//Using Bufio Scanner
	fmt.Println(readerFullText)

	if errr != nil {
		fmt.Println("Error")
	}
	fmt.Println(char)
	switch char {
	case 'A':
		fmt.Println(char)
		break
	case 'B':
		fmt.Println(char)
		break
	case 'C':
		fmt.Println(char)
		fallthrough
	case 'D':
		fmt.Println(char)
	default:
		fmt.Println(char)
	}

}
