package v0_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	v0 "github.com/Marcelixoo/learn-go-with-tests/json/v0"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := v0.NewInMemoryPlayerStore()
	server := v0.NewPlayerServer(store)
	player := "Pepper"

	// record 3 wins for player under test
	server.ServeHTTP(httptest.NewRecorder(), makeRecordWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), makeRecordWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), makeRecordWinRequest(player))

	response := httptest.NewRecorder()

	server.ServeHTTP(response, makeGetScoreRequest(player))

	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
}
