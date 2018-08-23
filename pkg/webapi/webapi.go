package webapi

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type webapi struct {
	Router *mux.Router
}

func NewWebapi(domain Domain) *webapi {
	a := webapi{}
	a.Router = mux.NewRouter()

	NewItemController(a.Router, domain)

	return &a
}

type Domain interface {
	GetItems(count int) ([]string, error)
	GetItem(s string) (string, error)
	SetItem(s string) (string, error)
}

func (a *webapi) Run(addr string) {
	log.Println("Server started")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
