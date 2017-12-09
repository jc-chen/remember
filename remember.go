package main

import (
	"encoding/json"
	"fmt"
	logging "github.com/op/go-logging"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

const app = "remember"

var log = logging.MustGetLogger(app)

var RMBFILE = os.Getenv("HOME") + "/.remember"

type Todo struct {
	Message   string    `json:message`
	Timestamp time.Time `json:timestamp`
}

type Remember struct {
	Todos []Todo `json:"todoList"`
}

var remember *Remember

func init() {
	// set log level
	logging.SetLevel(logging.INFO, app)

	content, err := ioutil.ReadFile(RMBFILE)
	if err != nil {
		content = initializeFile()
	}
	remember = &Remember{}
	json.Unmarshal(content, remember)
	log.Debugf("Done init: %+v", remember)
}

func initializeFile() []byte {
	log.Debug("initialize Remember")
	empty := []byte(`{"todoList":[]}`)
	write(empty)
	log.Debug("created init file")
	return empty
}

func (r *Remember) addTodo() {
	message := strings.Join(os.Args[1:], " ")
	todo := Todo{Message: message, Timestamp: time.Now()}
	r.Todos = append(r.Todos, todo)
	log.Debug("added new todo")
	writeToFile()
}

func (r *Remember) listTodo() {
	fmt.Println("List of Todos:")
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
		log.Error("Invalid command")
	}
	indexToDelete, err := strconv.Atoi(os.Args[2])
	checkErr(err)
	r.Todos = append(r.Todos[:indexToDelete-1], r.Todos[indexToDelete:]...)
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	switch os.Args[1] {
	case "ls":
		remember.listTodo()
	case "rm":
		remember.deleteTodo()
	default:
		remember.addTodo()
	}
	writeToFile()
}

func writeToFile() {
	res, err := json.Marshal(remember)
	checkErr(err)
	write(res)
}

func write(payload []byte) {
	ioutil.WriteFile(RMBFILE, payload, 0644)
}

func checkErr(err error) {
	if err != nil {
		log.Error(err)
	}
}
