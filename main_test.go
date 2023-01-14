package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-list-golang/models"
)

func InsertGetDeleteTodosTest(t *testing.T) {
	todoToInsert := models.Todo{Value: "Nick's todo", IsComplete: false}
	body, _ := json.Marshal(todoToInsert)
	_, err := http.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))

	if err != nil {
		fmt.Printf("error should be nil, but was: %v", err)
	}

	res := httptest.NewRecorder()
	fmt.Println(res)

}
