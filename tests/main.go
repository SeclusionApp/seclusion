package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	URL = "http://127.0.0.1:8080/v1/auth/register"
)

func main() {
	postSpam(URL)
}

func postSpam(url string) {
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		post(url, randomString(rand.Intn(1000)), randomString(rand.Intn(1000)))
	}
}

func post(url string, username string, password string) {

	fmt.Println("HTTP JSON POST URL:", url)

	var jsonData = []byte(`{
        "username": "` + username + `",
        "password": "` + password + `"
    }`)
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

}

func getSpam(url string) {
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()
	}
}

func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
