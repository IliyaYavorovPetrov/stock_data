package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	ticker := "TSLA"
	api := "e808bc63e1de4120a2690e7d4a447156"

	fmt.Println(ticker)
	fmt.Println(api)

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)
}
