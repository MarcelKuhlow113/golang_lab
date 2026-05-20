package auth

import (
    "fmt"
    "net/http"
)

func SetSessionCookie(w http.ResponseWriter, username string) {
    http.SetCookie(w, &http.Cookie{
        Name:     "session_user",
        Value:    username,
        Path:     "/",
        HttpOnly: true,
    })
}

func ClearSessionCookie(w http.ResponseWriter) {
    http.SetCookie(w, &http.Cookie{
        Name:   "session_user",
        Value:  "",
        Path:   "/",
        MaxAge: -1,
    })
}

func GetSessionUser(r *http.Request) (string, error) {
    cookie, err := r.Cookie("session_user")
    if err != nil {
        return "", err
    }
    if cookie.Value == "" {
        return "", fmt.Errorf("no session")
    }
    return cookie.Value, nil
}
