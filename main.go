package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/igorbednar/httpbend/sender"
)

func main() {

	url := flag.String("u", "", "server url")
	duration := flag.Int("d", 1000, "duration in seconds")
	rate := flag.Int("r", 10, "number of requests per second")

	flag.Parse()

	if len(strings.TrimSpace(*url)) == 0 {
		flag.PrintDefaults()
		return
	}

	sender := sender.HTTPSender{
		URL:           *url,
		RatePerSecond: *rate,
		Duration:      *duration,
	}

	result := sender.Start()
	fmt.Printf("Total number of requests sent: %d \n", result.TotalReqSent)
	fmt.Printf("Total number of errors : %d \n", result.NumOfErrors)
	fmt.Printf("Average response time is: %f sec \n", result.AvgReqTime)
}
