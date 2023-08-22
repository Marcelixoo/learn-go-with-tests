package v0

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerServer struct {
	store PlayerStore
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	return &PlayerServer{store}
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.recordWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) recordWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)

	w.WriteHeader(http.StatusAccepted)
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

func GetPlayerScore(player string) int {
	if player == "Pepper" {
		return 20
	}

	if player == "Floyd" {
		return 10
	}

	return 0
}
