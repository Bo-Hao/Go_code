package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func httpGet() {
	resp, err := http.Get("http://tw.yahoo.com")
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
}

func httpPost() {
	resp, err := http.Post("https://tw.yahoo.com/",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=test"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func main() {
	httpPost()

}
