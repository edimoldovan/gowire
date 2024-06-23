package handlers

import (
	"fmt"
	"net/http"
)

type Handlers struct{}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func (h *Handlers) PublicPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a public page!")
}
