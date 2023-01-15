package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-list-golang/models"

	"github.com/joho/godotenv"
)

func TestInsertGetDeleteTodos(t *testing.T) {
	godotenv.Load(".env")
	SetupDb()
	router := SetupRouter()

	todoToInsert := models.Todo{
		Id:         0,
		Value:      "Nick's todo",
		IsComplete: false,
	}

	body, _ := json.Marshal(todoToInsert)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	router.ServeHTTP(w, req)
	todo := decodeTodo(w)

	if w.Result().StatusCode != 200 {
		t.Errorf("StatusCode should be 200 but is: %v", w.Result().StatusCode)
	}

	if todo.Id == 0 {
		t.Errorf("todo id should be greater than 0 but was: %v", todo.Id)
	}

	if todo.Value != "Nick's todo" {
		t.Errorf("todo value should be: Nick's todo but was: %s", todo.Value)
	}

	if todo.IsComplete != false {
		t.Errorf("todo value should be: false but was %v", todo.IsComplete)
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/todos/%v", todo.Id), nil)
	router.ServeHTTP(w, req)
	todo = decodeTodo(w)

	if w.Result().StatusCode != 200 {
		t.Errorf("StatusCode should be 200 but is: %v", w.Result().StatusCode)
	}

	if todo.Id == 0 {
		t.Errorf("todo id should be greater than 0 but was: %v", todo.Id)
	}

	if todo.Value != "Nick's todo" {
		t.Errorf("todo value should be: Nick's todo but was: %s", todo.Value)
	}

	if todo.IsComplete != false {
		t.Errorf("todo value should be: false but was %v", todo.IsComplete)
	}

	todoToUpdate := models.Todo{
		Id:         todo.Id,
		Value:      "Nick's todo 1",
		IsComplete: true,
	}

	body, _ = json.Marshal(todoToUpdate)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPut, fmt.Sprintf("/todos/%v", todo.Id), bytes.NewReader(body))
	router.ServeHTTP(w, req)
	todo = decodeTodo(w)

	if w.Result().StatusCode != 200 {
		t.Errorf("StatusCode should be 200 but is: %v", w.Result().StatusCode)
	}

	if todo.Id != todoToUpdate.Id {
		t.Errorf("todo id should be the same as %v but was %v", todo.Id, todoToUpdate.Id)
	}

	if todo.Value != "Nick's todo 1" {
		t.Errorf("todo value should be: Nick's todo 1 but was: %s", todo.Value)
	}

	if !todo.IsComplete {
		t.Errorf("todo value should be: true but was %v", todo.IsComplete)
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/todos", nil)
	router.ServeHTTP(w, req)
	var todos []models.Todo
	json.NewDecoder(w.Body).Decode(&todos)

	if len(todos) == 0 {
		t.Errorf("todos should not be empty but is: %v", len(todos))
	}

	for _, val := range todos {
		if val.Id == 0 {
			t.Errorf("todo id should be greater than 0 but was: %v", todo.Id)
		}

		if val.Value != "Nick's todo 1" {
			t.Errorf("todo value should be: Nick's todo 1 but was: %s", todo.Value)
		}

		if !val.IsComplete {
			t.Errorf("todo value should be: false but was %v", todo.IsComplete)
		}
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodDelete, fmt.Sprintf("/todos/%v", todo.Id), nil)
	router.ServeHTTP(w, req)
	var truthy bool
	json.NewDecoder(w.Body).Decode(&truthy)

	if !truthy {
		t.Errorf("truthy should be: true but was %v", truthy)
	}
}

func TestGetBadRequest(t *testing.T) {
	godotenv.Load(".env")
	SetupDb()
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/todos/string", nil)
	router.ServeHTTP(w, req)

	if w.Result().StatusCode != 400 {
		t.Errorf("StatusCode should be 400 but is: %v", w.Result().StatusCode)
	}
}

func TestDeleteCantFindRecord(t *testing.T) {
	godotenv.Load(".env")
	SetupDb()
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/todos/100000", nil)
	router.ServeHTTP(w, req)

	if w.Result().StatusCode != 400 {
		t.Errorf("StatusCode should be 400 but is: %v", w.Result().StatusCode)
	}
}

func TestPutNonMatchingId(t *testing.T) {
	godotenv.Load(".env")
	SetupDb()
	router := SetupRouter()

	todoToUpdate := models.Todo{
		Id:         999999,
		Value:      "Nick's todo 1",
		IsComplete: true,
	}

	body, _ := json.Marshal(todoToUpdate)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/todos/1", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	if w.Result().StatusCode != 400 {
		t.Errorf("StatusCode should be 400 but is: %v", w.Result().StatusCode)
	}
}

func decodeTodo(w *httptest.ResponseRecorder) models.Todo {
	var todo models.Todo
	json.NewDecoder(w.Body).Decode(&todo)
	return todo
}
