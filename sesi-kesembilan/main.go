package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

func main() {

	PostData()


	
}

func GetData() {
	client := http.Client{}

	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil) 

	if err != nil {
		panic(err)
	}

	response, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var data map[string]interface{}

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		fmt.Println("error decode response body", err)
	}

	fmt.Println(data)
	fmt.Println("User ID => ", data["userId"])
	fmt.Println("ID => ", data["id"])
	fmt.Println("Title => ", data["title"])
	fmt.Println("Body =>", data["body"])


	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	sb := string(body)
	fmt.Println(sb)

	fmt.Println(reflect.TypeOf(sb))
}

func PostData() {
	var payload map[string]interface{} = map[string]interface{}{
		"title":"The Art of Zulkarnen",
		"body":"Zulkarnen's artisan",
		"userId": 1,
	} 


	requestJson, err := json.Marshal(&payload)

	if err != nil {
		panic(err)
	}

	// fmt.Println(string(requestJson))
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)

	if (err != nil) {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}