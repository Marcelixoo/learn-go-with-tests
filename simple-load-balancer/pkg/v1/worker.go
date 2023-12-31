package v1

import "fmt"

type Request struct {
	fn func() string // operation to perform.
	c  chan string   // channel to return the results.
}

func NewRequest(handler func() string, out chan string) Request {
	return Request{handler, out}
}

type Worker struct {
	id       string
	requests chan Request // work to do (buffered channel)
	pending  int          // count of pending tasks
	index    int          // index in the heap
	started  bool
}

func NewWorker(id string) *Worker {
	return &Worker{id, make(chan Request, 100), 0, 0, false}
}

func (w *Worker) Start(done chan *Worker) {
	appendID := func(res string) string {
		return fmt.Sprintf("%s [%s]", res, w.id)
	}

	if w.started {
		return
	}
	w.started = true

	for {
		req := <-w.requests // get request from balancer
		w.pending++
		req.c <- appendID(req.fn()) // send result to channel
		done <- w                   // signal work is done
		w.pending--
	}
}
