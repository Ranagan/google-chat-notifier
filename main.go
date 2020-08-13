package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

// RequestBody represents the body for a request sent to Google Chat
type RequestBody struct {
	Text string `json:"text"` // The text to post to chat
}

func makeRequest(url string, body RequestBody) *http.Response {
	var jsonData []byte
	jsonData, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println(err)
	}

	return resp
}

func main() {
	currentTime := time.Now().Unix()
	t := strconv.Itoa(int(currentTime))

	threadKey := flag.String("threadKey", t, "The thread to send the message to")
	chatUrl := flag.String("url", os.Getenv("CHATBOT_URL"), "The URL to send the chat message to. Defaults to the CHATBOT_URL env variable")
	message := flag.String("message", "Default message from chat notifier", "The message to send to the chat")

	flag.Parse()

	u, _ := url.Parse(*chatUrl)

	q := u.Query()
	q.Add("threadKey", *threadKey)
	u.RawQuery = q.Encode()

	body := RequestBody{
		Text: *message,
	}
	resp := makeRequest(u.String(), body)

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}
