package main

import (
	"bytes"
	"encoding/json"
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

	if w.Result().StatusCode != 200 {
		t.Errorf("StatusCode should be 200 but is: %v", w.Result().StatusCode)
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/todos", nil)
	router.ServeHTTP(w, req)

	var todos []models.Todo
	json.NewDecoder(w.Body).Decode(&todos)

	if len(todos) == 0 {
		t.Errorf("todos should not be empty but is: %v", len(todos))
	}
}
