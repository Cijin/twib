package handler

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
)

// render component should be used when redering items that are part of a page
// easy way to differentiate, is they usually don't render nav within
func renderComponent(w http.ResponseWriter, r *http.Request, component templ.Component) {
	err := component.Render(r.Context(), w)
	if err != nil {
		log.Println("Error rendering component: ", err)
	}
}

// helper, should never be used on it's own
func renderWithLayout(w http.ResponseWriter, r *http.Request, component templ.Component) {
	err := views.Layout(component).Render(r.Context(), w)
	if err != nil {
		log.Println("Error rendering page with layout: ", err)
	}
}

func renderPage(w http.ResponseWriter, r *http.Request, component templ.Component) {
	renderWithLayout(w, r, component)
}
