package v0_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	v0 "github.com/Marcelixoo/learn-go-with-tests/json/v0"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) ResetCounters() {
	s.winCalls = []string{}
}

func TestShowPlayerScores(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		[]string{},
	}
	server := v0.NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := makeGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := makeGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := makeGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		assertStatus(t, got, want)
	})
}

func TestRecordWins(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{},
		[]string{},
	}
	server := v0.NewPlayerServer(store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		store.ResetCounters()

		request := makeRecordWinRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
	})

	t.Run("it records a win", func(t *testing.T) {
		store.ResetCounters()

		player := "Pepper"

		request := makeRecordWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := 1
		got := len(store.winCalls)

		if got != want {
			t.Errorf("wrong number of calls to RecordWin, got %d want %d", got, want)
		}

		if store.winCalls[0] != player {
			t.Errorf(
				"did not record correct winner, got %q want %q",
				store.winCalls[0], player,
			)
		}
	})
}

func makeGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func makeRecordWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf(
			"did not get correct status code, got %d, want %d",
			got, want,
		)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf(
			"wrong response body, got %q want %q",
			got, want,
		)
	}
}
