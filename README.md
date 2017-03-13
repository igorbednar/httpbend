# httpbend
Simple http load tester written in go. Creates load on server by sending a constant number of http requests per second.

#TODO
Uses naive implementation where new goroutine is created per request.
A pool of goroutines should be used.
Statistics should be calculated and displayed
