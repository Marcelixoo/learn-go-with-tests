package v0

import (
	"fmt"
	"time"
)

type Work struct {
	x, y, z int
}

func Worker(in <-chan *Work, out chan<- *Work) {
	for w := range in {
		w.z = w.x * w.y
		time.Sleep(time.Duration(w.z) * time.Second)
		out <- w
	}
}

func Run(numOfWorkers, numOfJobs int) {
	in, out := make(chan *Work), make(chan *Work)

	for i := 0; i < numOfWorkers; i++ {
		go Worker(in, out)
	}

	go sendLotsOfWork(in, numOfJobs)
	receiveLotsOfResults(out)
}

func sendLotsOfWork(in chan *Work, numOfJobs int) {
	workload := make([]int, numOfJobs)

	for i := range workload {
		x, y, z := i, i, i
		in <- &Work{x, y, z}
	}
}
func receiveLotsOfResults(out <-chan *Work) {
	for result := range out {
		fmt.Printf("work processed %v", result)
	}
}
