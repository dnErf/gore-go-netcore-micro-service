package web

import (
	"bff/web/pages"
	"net/http"
)

func RenderHome(w http.ResponseWriter, r *http.Request) {
	pages.Home().Render(r.Context(), w)
}
