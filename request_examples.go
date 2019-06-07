package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func getTest() {
	log.Println("Get Test")

	url := "https://httpbin.org/get"

	response, err := http.Get(url)
	checkErr(err)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	checkErr(err)

	log.Println(string(body))

}

func postTest() {
	log.Println("Post Test")
	message := map[string]interface{}{
		"Hello": "world",
		"count": 1337,
		"nested": map[string]string{
			"Something": "Is nested!",
		},
	}

	url := "https://httpbin.org/post"
	byteJSON, err := json.Marshal(message)
	checkErr(err)

	payload := bytes.NewBuffer(byteJSON)
	header := "application/json"

	response, err := http.Post(url, header, payload)
	checkErr(err)

	var result map[string]interface{}

	json.NewDecoder(response.Body).Decode(&result)
	checkErr(err)

	log.Println(result)
	log.Println(result["data"])
}

func main() {
	getTest()
	postTest()
}
