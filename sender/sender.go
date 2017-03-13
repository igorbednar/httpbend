package sender

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//HTTPSender represents http load tester
type HTTPSender struct {
	URL           string
	Request       string
	ContentType   string
	RatePerSecond int
	NumOfRequests int
}

//Start starts load testing
func (s HTTPSender) Start() {
	timePerRequest := time.Duration(int64(time.Second) / int64(s.RatePerSecond))
	ch := make(chan string)
	done := make(chan bool)

	go waitForResponses(ch, done, s.NumOfRequests)

	for i := 0; i < s.NumOfRequests; i++ {
		start := time.Now()
		go sendRequest(s, ch)
		duration := time.Since(start)
		if duration < timePerRequest {
			time.Sleep(timePerRequest - duration)
		}
	}

	<-done
}

func sendRequest(s HTTPSender, ch chan<- string) {

	resp, err := http.Post(s.URL, s.ContentType, strings.NewReader(s.Request))
	if err != nil {
		ch <- fmt.Sprintf("while posting to %s: %v", s.URL, err)
		return
	}
	defer resp.Body.Close()
	_, err = io.Copy(ioutil.Discard, resp.Body)

	if err != nil {
		ch <- fmt.Sprintf("while reading response from %s: %v", s.URL, err)
		return
	}

	ch <- fmt.Sprintf("Received response from %s", s.URL)
}

func waitForResponses(ch chan string, done chan bool, numOfRequests int) {
	for i := 0; i < numOfRequests; i++ {
		fmt.Println(<-ch)
	}

	done <- true
}
