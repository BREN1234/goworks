package main

import (
	"fmt"
	"regexp"
	"strings"
)

func mainString() {
	//Slice of Byte
	str := "A regular expression is a sequence of characters which define a search pattern"
	str1 := "In Go language, strings are different from other languages like Java, C++, Python, etc. It is a sequence of variable-width characters where each and every character is represented by"
	m := regexp.MustCompile("e")
	dst := m.ReplaceAllString(str, "REPLACE")
	fmt.Println(dst)
	split1 := strings.Split(str1, " ")
	split2 := strings.Split(str1, ",")
	fmt.Println(split1, len(split1), split1[1])
	fmt.Println(split2, len(split2), split2[1])

	//SplitAfter(str,sep) []strings
	afterSplit1 := strings.SplitAfter(str, " ")
	afterSplit2 := strings.SplitAfter(str1, "I")
	fmt.Println(afterSplit1, len(afterSplit1), afterSplit1[1])
	fmt.Println(afterSplit2, len(afterSplit2), afterSplit2[1])

	//SplitAfterN(str,sep,m)
	splitN := strings.SplitAfterN(str1, ",", 3)
	splitN1 := strings.SplitAfterN(str, " ", 4)
	fmt.Println(splitN, len(splitN))
	fmt.Println(splitN1, len(splitN1))

	//MatchString(pattern string,s string)(bool , err)
	matched, err := regexp.MatchString("the", str1)
	fmt.Println(matched, err)
	//Finding string
	regExp := regexp.MustCompile(",")
	fmt.Println(regExp.FindString(str1))
	fmt.Println(regExp.FindAllString(str1, 10))
	//Index of strings
	regExp1 := regexp.MustCompile("of")
	fmt.Println(regExp1.FindStringIndex(str1))
	//slice of byte
	regExpByte, er := regexp.Match("str", []byte{'s', 'o', 'r', 'i', 't', 'g'})
	fmt.Println(regExpByte, er, []byte("asdfasdfasdf"))

}
