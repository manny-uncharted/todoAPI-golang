package main

import (
	"log"
	"net/http"
	"os"
)

var testingMode bool

func init() {
	testingMode = os.Getenv("TESTING_MODE") == "true"
}

func main() {

	initDatabase()
	defer Db.Close()

	http.HandleFunc("/todos", loggingMiddleware(http.HandlerFunc(getTodos)))
	http.HandleFunc("/todo", loggingMiddleware(http.HandlerFunc(createTodo)))
	http.Handle("/todo/update", loggingMiddleware(http.HandlerFunc(updateTodo)))
	http.Handle("/todo/delete", loggingMiddleware(http.HandlerFunc(deleteTodo)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
