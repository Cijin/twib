package handlers

import (
	"net/http"

	"twib/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	renderComponent(w, r, home.Index())
}
