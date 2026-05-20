package login

import (
	"net/http"

	"golab/app"
	"golab/components/auth"
	"golab/components/database"
)

func RegisterRoutes(mux *http.ServeMux, application *app.App) {
	mux.HandleFunc("/login", loginHandler(application))
}

func loginHandler(application *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := app.ViewData{}
		if msg := r.URL.Query().Get("msg"); msg == "registered" {
			data.Success = "Registration successful! Please log in."
		}

		if r.Method != http.MethodPost {
			application.Render(w, "login.html", data)
			return
		}

		if err := r.ParseForm(); err != nil {
			data.Error = "Unable to read form data"
			application.Render(w, "login.html", data)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" || password == "" {
			data.Error = "Please enter both username and password"
			application.Render(w, "login.html", data)
			return
		}

		if !database.UserNameExists(username) {
			data.Error = "Invalid username or password"
			application.Render(w, "login.html", data)
			return
		}

		if !database.VerifyLogin(username, password) {
			data.Error = "Invalid username or password"
			application.Render(w, "login.html", data)
			return
		}

		if !database.VerifyLogin(username, password) {
			data.Error = "Invalid username or password"
			application.Render(w, "login.html", data)
			return
		}

		auth.SetSessionCookie(w, username)
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
	}
}
