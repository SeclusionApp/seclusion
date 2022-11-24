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

var usernames = []string{}
var passwords = []string{}

const (
	URL = "http://127.0.0.1:8080/v1/"
)

func main() {
	postSpamRegister(URL + "auth/")
	getSpamMessages(URL + "messages/")
	postSpamChannels()
}

func postSpamChannels() {
	var channel_url = URL + "channels/"
	for {
		for i := 0; i < 100; i++ {
			channel_name := randomString(10)
			channel_desc := randomString(50)
			channel := fmt.Sprintf(`{"name":"%s"`)
			post(channel_url, channel_name, channel_desc)
		}
		time.Sleep(1 * time.Second)
	}
}

func getSpamMessages(url string) {
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		login(url + "login")
		postMessage(url + "message/")
		getMessages(url + "messages/")
	}
}

func postMessage(url string) {
	fmt.Println("HTTP JSON POST URL:", url+"message")

	var randInt = rand.Intn(len(usernames))
	var jsonData = []byte(`{
		"channel_id": "1",
		"user_id": "1",
		"message": "` + usernames[randInt] + `"
	}`)

	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		log.Fatal(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
}

func getMessages(url string) {
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		getMessages(url)
	}
}

func login(url string) {
	fmt.Println("HTTP JSON POST URL:", url+"login")

	var randInt = rand.Intn(len(usernames))
	var jsonData = []byte(`{
		"username": "` + usernames[randInt] + `",
		"password": "` + passwords[randInt] + `"
	}`)

	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		log.Fatal(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
}

func postSpamRegister(url string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		post(url+"register", randomString(rand.Intn(1000)), randomString(rand.Intn(1000)))
	}
}

func post(url string, username string, password string) {

	fmt.Println("HTTP JSON POST URL:", url+"register")

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

	if response.StatusCode == 200 {
		usernames = append(usernames, username)
		passwords = append(passwords, password)
	}
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
