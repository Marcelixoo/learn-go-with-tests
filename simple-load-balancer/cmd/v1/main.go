package main

import (
	"fmt"
	"math/rand"
	"time"

	v1 "github.com/Marcelixoo/learn-go-with-tests/simple-load-balancer/pkg/v1"
)

func main() {
	work := make(chan v1.Request)
	response := make(chan string)

	poolOfWorkers := v1.Pool{
		v1.NewWorker("#1"),
		v1.NewWorker("#2"),
		v1.NewWorker("#3"),
		v1.NewWorker("#4"),
	}

	numOfWorkers := len(poolOfWorkers)
	go StartRequester(work, response, numOfWorkers)

	balancer := v1.NewBalancer(poolOfWorkers)
	go balancer.Balance(work)

	fmt.Println("processing incoming requests")
	for v := range response {
		fmt.Printf("got response %q\n", v)
	}
}

func StartRequester(work chan<- v1.Request, response chan<- string, numOfWorkers int) {
	c := make(chan string)

	for i := 0; i < numOfWorkers; i++ {
		go func(numOfWorkers int) {
			for {
				simulateLoadFor(numOfWorkers)

				work <- v1.NewRequest(workerFn, c)
				result := <-c
				response <- result
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
