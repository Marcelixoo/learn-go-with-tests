package v1

type Pool []*Worker

// Pool is an implementation of the Heap interface.
// It prioritizes items in the list based on load.
func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p Pool) Len() int {
	return len(p)
}

func (p Pool) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

func (p *Pool) Push(x any) {
	n := len(*p)
	worker := x.(*Worker)
	worker.index = n
	*p = append(*p, worker)
}

func (p *Pool) Pop() any {
	old := *p
	n := len(old)
	worker := old[n-1]
	old[n-1] = nil    // avoid memory leak
	worker.index = -1 // for safety
	*p = old[0 : n-1]
	return worker
}
