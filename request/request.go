package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	request, _ := http.NewRequest("GET", "http://google.co.jp", nil)

	client := new(http.Client)

	response, _ := client.Do(request)

	defer response.Body.Close()

	byteArray, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(byteArray))
}
