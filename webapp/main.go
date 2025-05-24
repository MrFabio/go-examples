package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type ToDoList struct {
	ToDoCount int
	ToDos     []string
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func write(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	errorCheck(err)
}

func getStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	defer file.Close()
	errorCheck(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	errorCheck(scanner.Err())
	return lines, nil
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello Internet!")
}

func toDoHandler(writer http.ResponseWriter, request *http.Request) {
	lines, err := getStrings("todos.txt")
	errorCheck(err)
	fmt.Printf("%#v\n", lines)
	tmpl, err := template.ParseFiles("html/view.html")
	errorCheck(err)

	todos := ToDoList{
		ToDoCount: len(lines),
		ToDos:     lines,
	}
	tmpl.Execute(writer, todos)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("html/new.html")
	errorCheck(err)
	tmpl.Execute(writer, nil)
}

func createTodoHandler(writer http.ResponseWriter, request *http.Request) {
	todo := request.FormValue("todo")
	opts := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("todos.txt", opts, os.FileMode(0600))
	errorCheck(err)
	_, err = fmt.Fprintln(file, todo)
	errorCheck(err)
	http.Redirect(writer, request, "/todos", http.StatusSeeOther)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", englishHandler)
	mux.HandleFunc("/todos", toDoHandler)
	mux.HandleFunc("/todos/new", newHandler)
	mux.HandleFunc("/todos/create", createTodoHandler)

	server := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	log.Println("Starting server on :4000")
	err := server.ListenAndServe()
	errorCheck(err)
}
