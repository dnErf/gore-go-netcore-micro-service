package views

import (
	"net/http"

	"gore/views/pages"
)

func RenderHome(w http.ResponseWriter, r *http.Request) {
	pages.Home().Render(r.Context(), w)
}
