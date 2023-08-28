package v1

import (
	"container/heap"
)

type Balancer struct {
	pool Pool
	done chan *Worker
}

func NewBalancer(pool Pool) *Balancer {
	b := Balancer{pool, make(chan *Worker)}

	for _, worker := range pool {
		go worker.Start(b.done)
	}

	return &b
}

func (b *Balancer) Balance(work chan Request) {
	for {
		select {
		case req := <-work:
			b.dispatch(req)
		case w := <-b.done:
			b.completed(w)
		}
	}
}

// dispatches request to the least loaded worker
func (b *Balancer) dispatch(request Request) {
	w := heap.Pop(&b.pool).(*Worker)
	w.requests <- request
	heap.Push(&b.pool, w)
}

func (b *Balancer) completed(w *Worker) {
	heap.Remove(&b.pool, w.index)
	heap.Push(&b.pool, w)
}
