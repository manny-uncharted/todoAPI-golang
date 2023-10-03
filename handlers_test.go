package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTodos(t *testing.T) {
	initDatabase()
	req, _ := http.NewRequest("GET", "/todos", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getTodos)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	defer Db.Close()
}

func TestCreateTodo(t *testing.T) {
	initDatabase()
	todo := Todo{
		Title: "Test Todo",
		Done:  false,
	}
	data, _ := json.Marshal(todo)
	req, _ := http.NewRequest("POST", "/todo", bytes.NewBuffer(data))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createTodo)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var responseTodo Todo
	_ = json.Unmarshal(rr.Body.Bytes(), &responseTodo)
	assert.Equal(t, todo.Title, responseTodo.Title)
	defer Db.Close()
}

// Test for updating a Todo
func TestUpdateTodo(t *testing.T) {
	initDatabase()

	todo := Todo{
		ID:    1, // Assuming this ID exists in the database
		Title: "Updated Todo",
		Done:  true,
	}
	data, _ := json.Marshal(todo)
	req, _ := http.NewRequest("PUT", "/todo/update", bytes.NewBuffer(data))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateTodo)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var responseTodo Todo
	_ = json.Unmarshal(rr.Body.Bytes(), &responseTodo)
	assert.Equal(t, todo.Title, responseTodo.Title)
	assert.Equal(t, todo.Done, responseTodo.Done)
	defer Db.Close()
}

// Test for deleting a Todo
func TestDeleteTodo(t *testing.T) {
	initDatabase()

	req, _ := http.NewRequest("DELETE", "/todo/delete?id=1", nil) // Assuming ID 1 exists in the database
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteTodo)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	defer Db.Close()
}
