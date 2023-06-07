package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Use this code to constantly ask for a key while you parallely perform other write
// operations, this is to test concurrency
func main() {

	url := "http://localhost:8000/read/"
	method := "POST"

	payload := strings.NewReader(`{
    "key": "user123"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	for {
		time.Sleep(time.Second)
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body))
	}
}
