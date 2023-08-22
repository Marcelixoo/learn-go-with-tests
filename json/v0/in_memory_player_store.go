package v0

type InMemoryPlayerStore struct {
	scores map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.scores[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.scores[name]++
}
