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
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
