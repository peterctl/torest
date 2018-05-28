package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// This App struct wraps the repository and router to be
// able to send then into the routes whenever needed.
type App struct {
	Repo   ToDoRepo
	Router *mux.Router
}

// App.HandleFunc wraps mux.Router.HandleFunc to inject itself into
// the route handler.
func (a *App) HandleFunc(
	route string,
	handler func(a *App, w http.ResponseWriter, r *http.Request),
) *mux.Route {
	return a.Router.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		handler(a, w, r)
	})
}

func (a *App) ListenAndServe(addr string) error {
	handler := cors.AllowAll().Handler(a.Router)
	return http.ListenAndServe(addr, handler)
}

func NewApp(repo ToDoRepo) *App {
	app := &App{
		Repo:   repo,
		Router: mux.NewRouter(),
	}
	configureRoutes(app)
	return app
}

func configureRoutes(a *App) {
	a.HandleFunc("/api/todos", TodoIndex).Methods("GET")
	a.HandleFunc("/api/todos", CreateTodo).Methods("POST")
	a.HandleFunc("/api/todos/{id}", SingleTodo).Methods("GET")
	a.HandleFunc("/api/todos/{id}", UpdateTodo).Methods("PUT", "PATCH")
	a.HandleFunc("/api/todos/{id}", DeleteTodo).Methods("DELETE")
}
