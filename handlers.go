package main

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"golab/app"
	"golab/components/auth"
	"golab/components/database"
)

func redirectToLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func registerHandler(application *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := app.ViewData{}

		if r.Method != http.MethodPost {
			application.Render(w, "register.html", data)
			return
		}

		if err := r.ParseForm(); err != nil {
			data.Error = "Unable to read form data"
			application.Render(w, "register.html", data)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")
		confirm := r.FormValue("confirm")
		data.Username = username

		if username == "" || password == "" || confirm == "" {
			data.Error = "All fields are required"
			application.Render(w, "register.html", data)
			return
		}

		if password != confirm {
			data.Error = "Passwords do not match"
			application.Render(w, "register.html", data)
			return
		}

		if application.Store.Exists(username) {
			data.Error = "Username already exists"
			application.Render(w, "register.html", data)
			return
		}

		if database.UserNameExists(username) {
			data.Error = "Username already exists"
			application.Render(w, "register.html", data)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			data.Error = "Unable to register user"
			application.Render(w, "register.html", data)
			return
		}
		fmt.Print("Save Data into DB")
		database.SaveNewUser(username, password)
		application.Store.SaveHash(username, hash)
		http.Redirect(w, r, "/login?msg=registered", http.StatusSeeOther)
	}
}

func welcomeHandler(application *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, err := auth.GetSessionUser(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		application.Render(w, "welcome.html", app.ViewData{Username: username})
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	auth.ClearSessionCookie(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
