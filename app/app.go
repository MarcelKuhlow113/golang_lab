package app

import (
    "html/template"
    "sync"
)

type UserStore struct {
    sync.RWMutex
    users map[string][]byte
}

type ViewData struct {
    Error    string
    Success  string
    Username string
}

type App struct {
    Store *UserStore
    Tmpl  *template.Template
}

func NewUserStore() *UserStore {
    return &UserStore{users: map[string][]byte{}}
}

func (s *UserStore) Exists(username string) bool {
    s.RLock()
    defer s.RUnlock()
    _, ok := s.users[username]
    return ok
}

func (s *UserStore) SaveHash(username string, hash []byte) {
    s.Lock()
    defer s.Unlock()
    s.users[username] = hash
}

func (s *UserStore) FetchHash(username string) ([]byte, bool) {
    s.RLock()
    defer s.RUnlock()
    hash, ok := s.users[username]
    return hash, ok
}

func New() *App {
    return &App{
        Store: NewUserStore(),
        Tmpl:  template.Must(template.ParseGlob("templates/*.html")),
    }
}
