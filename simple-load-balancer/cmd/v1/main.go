package main

import (
	"fmt"
	"math/rand"
	"time"

	v1 "github.com/Marcelixoo/learn-go-with-tests/simple-load-balancer/pkg/v1"
)

func main() {
	channelOfRequests := make(chan v1.Request)
	channelOfResponses := make(chan string)

	poolOfWorkers := v1.Pool{
		v1.NewWorker("#1"),
		v1.NewWorker("#2"),
		v1.NewWorker("#3"),
		v1.NewWorker("#4"),
	}

	balancer := v1.NewBalancer(poolOfWorkers)
	go balancer.Balance(channelOfRequests)

	go StartRequester(channelOfRequests, channelOfResponses, poolOfWorkers.Len())

	fmt.Println("processing incoming requests")
	for res := range channelOfResponses {
		fmt.Printf("got response %q\n", res)
	}
}

func StartRequester(chanOfRequests chan<- v1.Request, chanOfResponses chan<- string, numOfWorkers int) {
	c := make(chan string)

	for i := 0; i < numOfWorkers; i++ {
		go func(numOfWorkers int) {
			for {
				simulateLoadFor(numOfWorkers)

				chanOfRequests <- v1.NewRequest(workerFn, c)
				result := <-c
				chanOfResponses <- result
			}
		}(numOfWorkers)
	}
}

func simulateLoadFor(numOfWorkers int) {
	seed := rand.Int63n(int64(numOfWorkers * 2))
	time.Sleep(time.Duration(seed) * time.Second)
}

func workerFn() string {
	return "doing some work"
}
