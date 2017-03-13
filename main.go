package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/igorbednar/httpbend/sender"
)

func main() {

	url := flag.String("u", "", "server url")
	filepath := flag.String("f", "", "file with request content")
	reqnum := flag.Int("n", 1000, "total number of requests")
	rate := flag.Int("r", 10, "number of requests per second")

	flag.Parse()

	if len(strings.TrimSpace(*url)) == 0 {
		flag.PrintDefaults()
		return
	}

	if len(strings.TrimSpace(*filepath)) == 0 {
		flag.PrintDefaults()
		return
	}

	file, err := ioutil.ReadFile(*filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}

	sender := &sender.HTTPSender{
		URL:           "http://localhost:8000",
		Request:       string(file),
		ContentType:   "text/xml",
		RatePerSecond: *rate,
		NumOfRequests: *reqnum,
	}

	sender.Start()
}
