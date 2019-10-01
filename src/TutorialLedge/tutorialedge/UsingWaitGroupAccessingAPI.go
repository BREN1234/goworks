package tutorialedge

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var Urls []string = []string{
	"https://google.com",
	"https://tutorialedge.net",
	"https://twitter.com",
}

func fetch(url string, waitGroup *sync.WaitGroup) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	waitGroup.Done()
	fmt.Println(resp.Status)
	return resp.Status, nil
}
func homePage(res http.ResponseWriter, req *http.Request) {
	fmt.Println("HomePage")
	var waitGroup sync.WaitGroup
	for _, url := range Urls {
		waitGroup.Add(1)
		go fetch(url, &waitGroup)
	}
	waitGroup.Wait()
	fmt.Println("Returning Response")
	fmt.Fprintf(res, "Responses")
}
func HandleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
