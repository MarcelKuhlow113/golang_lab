package app

import (
    "net/http"
)

func (a *App) Render(w http.ResponseWriter, name string, data ViewData) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    if err := a.Tmpl.ExecuteTemplate(w, name, data); err != nil {
        http.Error(w, "Server error: "+err.Error(), http.StatusInternalServerError)
    }
}
