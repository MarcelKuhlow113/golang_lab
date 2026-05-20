package main

import (
	"fmt"
	"golab/app"
	"golab/components/login"
	"log"
	"net/http"
)

func main() {
	application := app.New()
	mux := http.NewServeMux()

	login.RegisterRoutes(mux, application)
	mux.HandleFunc("/", redirectToLogin)
	mux.HandleFunc("/register", registerHandler(application))
	mux.HandleFunc("/welcome", welcomeHandler(application))
	mux.HandleFunc("/logout", logoutHandler)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
