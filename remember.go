package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

type Todo struct {
	Message   string    `json:message`
	Timestamp time.Time `json:timestamp`
}

type Remember struct {
	Todos []Todo `json:"todoList"`
}

func (r *Remember) addTodo() {
	message := strings.Join(os.Args[1:], " ")
	todo := Todo{Message: message, Timestamp: time.Now()}
	r.Todos = append(r.Todos, todo)
	log.Debug("added new todo")
	r.writeToFile()
}

// TODO have different ways of printing
func (r *Remember) listTodo() {
	fmt.Println("Your list of Todos:")
	// TODO have max line length based on terminal width and break line
	// so long todos don't break printing
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for i, todo := range r.Todos {
		fmt.Fprintf(w, " %d.\t%s\t\t%s\n", i+1,
			todo.Message,
			todo.Timestamp.Format("(15:04, Mon, Jan 2, 2006)"))
	}
	w.Flush()
}

func (r *Remember) deleteTodo() {
	if len(os.Args) != 3 {
		log.Error("Invalid command: missing arguments")
	}
	indexToDelete, err := strconv.Atoi(os.Args[2])
	checkErr(err)
	r.Todos = append(r.Todos[:indexToDelete-1], r.Todos[indexToDelete:]...)
	r.writeToFile()
}

func (r *Remember) writeToFile() {
	res, err := json.Marshal(r)
	checkErr(err)
	write(res)
}
