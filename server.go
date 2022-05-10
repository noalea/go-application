package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	SavePlayerScore(name string)
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.LeagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.PlayerHandler))

	p.Handler = router

	return p
}

func (p *PlayerServer) LeagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) PlayerHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.saveScore(w, player)
	case http.MethodGet:
		p.getScore(w, player)
	}
}

func (p *PlayerServer) getScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

func (p *PlayerServer) saveScore(w http.ResponseWriter, player string) {
	p.store.SavePlayerScore(player)
	w.WriteHeader(http.StatusAccepted)
}

func GetPlayerScore(path string) string {
	if path == "Pepper" {
		return "20"
	}
	if path == "Floyd" {
		return "10"
	}

	return ""
}
