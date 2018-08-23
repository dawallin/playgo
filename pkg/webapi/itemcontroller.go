package webapi

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type itemController struct {
	domain Domain
}

func NewItemController(router *mux.Router, domain Domain) *itemController {
	c := itemController{}

	c.domain = domain

	router.HandleFunc("/items", c.getItems).Methods("GET")
	router.HandleFunc("/item/{s}", c.setItem).Methods("POST")
	router.HandleFunc("/item/{id}", c.getItem).Methods("GET")
	//c.Router.HandleFunc("/item/{id:[0-9]+}", c.updateProduct).Methods("PUT")
	//c.Router.HandleFunc("/item/{id:[0-9]+}", c.deleteProduct).Methods("DELETE")

	return &c
}

func (c *itemController) getItems(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))

	if count > 10 || count < 1 {
		count = 10
	}

	products, err := c.domain.GetItems(count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (c *itemController) getItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		respondWithError(w, http.StatusInternalServerError, "Invalid id")
		return
	}

	products, err := c.domain.GetItem(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (c *itemController) setItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s := vars["s"]
	log.Println(s)

	if s == "" {
		respondWithError(w, http.StatusInternalServerError, "Invalid item")
		return
	}

	id, err := c.domain.SetItem(s)
	if err != nil {
		log.Println("Error: " + err.Error())

		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println(id)

	respondWithJSON(w, http.StatusOK, id)
}
