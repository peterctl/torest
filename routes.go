package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func TodoIndex(a *App, w http.ResponseWriter, r *http.Request) {
	todos, err := a.Repo.All()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(todos)
	}
}

func CreateTodo(a *App, w http.ResponseWriter, r *http.Request) {
	var data struct {
		Name *string `json:"name"`
	}
	json.NewDecoder(r.Body).Decode(&data)

	if data.Name == nil {
		var res struct {
			Errors struct {
				Name string `json:"name"`
			} `json:"errors"`
		}
		res.Errors.Name = "Name is required"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&res)
		return
	}

	todo := NewToDo(*data.Name, false)
	err := a.Repo.Save(todo)
	if err != nil {
		var res struct {
			Error string `json:"name"`
		}
		res.Error = err.Error()
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(&res)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func SingleTodo(a *App, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todo, err := a.Repo.Find(params["id"])
	if err != nil {
		var res struct {
			Error string `json:"name"`
		}
		res.Error = err.Error()
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(&res)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(a *App, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todo, err := a.Repo.Find(params["id"])
	if err != nil {
		var res struct {
			Error string `json:"name"`
		}
		res.Error = err.Error()
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(&res)
		return
	}

	var data map[string]*json.RawMessage
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil || (data["name"] == nil && data["completed"] == nil) {
		var res struct {
			Error string `json:"error"`
		}
		res.Error = "Empty request data."
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&res)
		return
	}

	if data["name"] != nil {
		json.Unmarshal(*data["name"], &todo.Name)
	}
	if data["completed"] != nil {
		json.Unmarshal(*data["completed"], &todo.Completed)
	}

	err = a.Repo.Save(todo)
	if err != nil {
		var res struct {
			Error string `json:"error"`
		}
		res.Error = err.Error()
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(res)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(a *App, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todo, err := a.Repo.Find(params["id"])
	if err != nil {
		var res struct {
			Error string `json:"error"`
		}
		res.Error = err.Error()
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(&res)
		return
	}

	err = a.Repo.Delete(todo)
	if err != nil {
		var res struct {
			Error string `json:"error"`
		}
		res.Error = err.Error()
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(res)
		return
	}

	json.NewEncoder(w).Encode(todo)
}
