package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:'social'`
}

type Social struct {
	Facebook string `json:"Facebook"`
	Twitter  string `json:"twitter"`
}

func mainJson() {

	//Json file
	//open ing the file
	file, err := os.Open("src/Files/User.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opend")
	defer file.Close()
	//UnMarshalling to json
	byteValue, _ := ioutil.ReadAll(file)
	fmt.Println(string(byteValue))
	var users Users
	json.Unmarshal(byteValue, &users)
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("name: ", users.Users[i].Name)
		fmt.Println("Age: ", strconv.Itoa(users.Users[i].Age))
		fmt.Println("Tyep: ", users.Users[i].Type)
		fmt.Println("Facebook: ", users.Users[i].Social.Facebook)
		fmt.Println("Twitter: ", users.Users[i].Social.Twitter)
	}

}
