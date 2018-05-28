package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	// "."
)

func app() (a *App) {
	a = NewApp(NewMapRepo())
	return
}

func processRequest(a *App, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func TestTodoIndex(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/todos", nil)
	res := processRequest(app(), req)

	// t.Log(res.Body.String())

	if res.Code != 200 {
		t.Errorf("Expected code to be %d but got %d.", 200, res.Code)
	}

	var todos []ToDo
	json.NewDecoder(res.Body).Decode(&todos)
	if len(todos) != 0 {
		t.Errorf("Expected len(todos) to be %d but got %d.", 0, len(todos))
	}
}

func TestAddTodoShouldSucceed(t *testing.T) {
	a := app()
	req, _ := http.NewRequest("POST", "/api/todos",
		strings.NewReader(`{ "name": "Test todo." }`))
	res := processRequest(a, req)

	// t.Log(res.Body.String())

	if res.Code != 200 {
		t.Errorf("Expected code to be %d but got %d.", 200, res.Code)
	}

	var todo ToDo
	json.NewDecoder(res.Body).Decode(&todo)
	if todo.Name != "Test todo." {
		t.Errorf("Expected name to be '%s' but got '%s'.", "Test todo.", todo.Name)
	}
	if todo.Completed != false {
		t.Errorf("Expected completed to be %t but got %t.", false, todo.Completed)
	}

	todos, _ := a.Repo.All()
	if len(todos) != 1 {
		t.Errorf("Expected len(todos) to be %d but got %d.", 1, len(todos))
	}
}

func TestAddTodoShouldFail(t *testing.T) {
	a := app()
	req, _ := http.NewRequest("POST", "/api/todos",
		strings.NewReader(""))
	res := processRequest(a, req)

	// t.Log(res.Body.String())

	if res.Code != 400 {
		t.Errorf("Expected code to be %d but got %d.", 400, res.Code)
	}
}

func TestSingleTodoShouldSucceed(t *testing.T) {
	a := app()
	a.Repo.Save(NewToDo("Test todo.", false))

	todos, _ := a.Repo.All()
	exp := todos[0]

	url := fmt.Sprintf("/api/todos/%s", exp.Id)
	req, _ := http.NewRequest("GET", url, nil)
	res := processRequest(a, req)

	// t.Log(res.Body.String())

	// Test the response code
	if res.Code != 200 {
		t.Errorf("Expected code to be %d but got %d.", 200, res.Code)
	}

	// Test that the returned to-do is the correct one
	var act ToDo
	json.NewDecoder(res.Body).Decode(&act)
	if exp.Id != act.Id {
		t.Errorf("Expected Id to be '%s' but got '%s'.", exp.Id, act.Id)
	}
	if exp.Name != act.Name {
		t.Errorf("Expected Name to be '%s' but got '%s'.", exp.Name, act.Name)
	}
	if exp.Completed != act.Completed {
		t.Errorf("Expected Completed to be %t but got %t.", exp.Completed, act.Completed)
	}
}

func TestSingleTodoShouldFail(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/todos/_", nil)
	res := processRequest(app(), req)

	// t.Log(res.Body.String())

	if res.Code != 404 {
		t.Errorf("Expected code to be %d but got %d", 404, res.Code)
	}
}

func TestUpdateTodoShouldSucceed(t *testing.T) {
	a := app()
	a.Repo.Save(NewToDo("Test todo.", false))

	todos, _ := a.Repo.All()
	todo := todos[0]

	url := fmt.Sprintf("/api/todos/%s", todo.Id)
	req, _ := http.NewRequest("PUT", url,
		strings.NewReader(`{ "name": "New name", "completed": true }`))
	res := processRequest(a, req)

	// t.Log(res.Body.String())

	// Test the response code
	if res.Code != 200 {
		t.Errorf("Expected code to be %d but got %d.", 200, res.Code)
	}

	// Test that the fields were updated correctly.
	json.NewDecoder(res.Body).Decode(&todo)
	if todo.Name != "New name" {
		t.Errorf("Expected Name to be '%s' but got '%s'.", "New name", todo.Name)
	}
	if todo.Completed != true {
		t.Errorf("Expected Completed to be %t but got %t.", true, todo.Completed)
	}
}

func TestUpdateTodoShouldFailWith404(t *testing.T) {
	req, _ := http.NewRequest("PUT", "/api/todos/_",
		strings.NewReader(""))
	res := processRequest(app(), req)

	// t.Log(res.Body.String())

	// Test the response code
	if res.Code != 404 {
		t.Errorf("Expected code to be %d but got %d.", 400, res.Code)
	}
}

func TestUpdateTodoShouldFailWith400(t *testing.T) {
	a := app()
	a.Repo.Save(NewToDo("Test todo.", false))

	todos, _ := a.Repo.All()
	todo := todos[0]

	url := fmt.Sprintf("/api/todos/%s", todo.Id)
	req, _ := http.NewRequest("PUT", url,
		strings.NewReader(""))
	res := processRequest(a, req)

	// t.Log(res.Body.String())

	// Test the response code
	if res.Code != 400 {
		t.Errorf("Expected code to be %d but got %d.", 400, res.Code)
	}
}

func TestDeleteTodoShouldSucceed(t *testing.T) {
	a := app()
	a.Repo.Save(NewToDo("Test todo.", false))

	todos, _ := a.Repo.All()
	exp := todos[0]
	ec := len(todos) - 1 // Expected count

	url := fmt.Sprintf("/api/todos/%s", exp.Id)
	req, _ := http.NewRequest("DELETE", url,
		strings.NewReader(`{ "completed": true }`))
	res := processRequest(a, req)

	// t.Log(res.Body.String())

	if res.Code != 200 {
		t.Errorf("Expected code to be %d but got %d.", 200, res.Code)
	}

	todos, _ = a.Repo.All()
	ac := len(todos) // Actual count

	if ec != ac {
		t.Errorf("Expected count to be %d but got %d.", ac, ec)
	}

	var act ToDo
	json.NewDecoder(res.Body).Decode(&act)
	if act.Id != exp.Id {
		t.Errorf("Expected Id to be '%s' but got '%s'.", exp.Id, act.Id)
	}
	if act.Name != exp.Name {
		t.Errorf("Expected Name to be '%s' but got '%s'.", exp.Name, act.Name)
	}
	if act.Completed != exp.Completed {
		t.Errorf("Expected Completed to be %t but got %t.", exp.Completed, act.Completed)
	}
}

func TestDeleteTodoShouldFail(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/api/todos/_", nil)
	res := processRequest(app(), req)

	if res.Code != 404 {
		t.Errorf("Expected code to be %d but got %d", 404, res.Code)
	}
}
