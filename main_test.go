package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-team/models"

	"github.com/stretchr/testify/assert"
)

func TestCleanupBefore(t *testing.T) {
	models.GetDB().Exec("TRUNCATE TABLE players RESTART IDENTITY")
	models.GetDB().Exec("TRUNCATE TABLE teams RESTART IDENTITY")
}

func TestCreateTeam(t *testing.T) {

	var jsonStr = []byte(`{"name":"Test Create"}`)

	r, _ := http.NewRequest("POST", "/api/teams", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	appRoute().ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "\"status\":true")
}

func TestGetTeams(t *testing.T) {

	r, _ := http.NewRequest("GET", "/api/teams", nil)
	w := httptest.NewRecorder()

	appRoute().ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "\"status\":true")
}

func TestFindTeam(t *testing.T) {

	r, _ := http.NewRequest("GET", "/api/teams/1", nil)
	w := httptest.NewRecorder()

	appRoute().ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "\"status\":true")
}

func TestUpdateTeam(t *testing.T) {

	var jsonStr = []byte(`{"name":"Test Update"}`)

	r, _ := http.NewRequest("PUT", "/api/teams/1", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	appRoute().ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "\"status\":true")
}

func TestDeleteTeam(t *testing.T) {

	r, _ := http.NewRequest("DELETE", "/api/teams/1", nil)
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	appRoute().ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "\"status\":true")
}

func TestCreatePlayer(t *testing.T) {

	var jsonStr = []byte(`{"name":"Test Player Name", "position":"Test Position", "team_id":1}`)

	r, _ := http.NewRequest("POST", "/api/players", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	appRoute().ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "\"status\":true")
}

func TestGetPlayers(t *testing.T) {

	r, _ := http.NewRequest("GET", "/api/players", nil)
	w := httptest.NewRecorder()

	appRoute().ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "\"status\":true")
}

func TestFindPlayer(t *testing.T) {

	r, _ := http.NewRequest("GET", "/api/players/1", nil)
	w := httptest.NewRecorder()

	appRoute().ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "\"status\":true")
}

func TestUpdatePlayer(t *testing.T) {

	var jsonStr = []byte(`{"name":"Test Update Player Name", "position":"Test Update Position", "team_id":1}`)

	r, _ := http.NewRequest("PUT", "/api/players/1", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	appRoute().ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "\"status\":true")
}

func TestDeletePlayer(t *testing.T) {

	r, _ := http.NewRequest("DELETE", "/api/players/1", nil)
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	appRoute().ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "\"status\":true")
}
