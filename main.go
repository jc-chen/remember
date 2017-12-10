package main

import (
	"encoding/json"
	logging "github.com/op/go-logging"
	"io/ioutil"
	"os"
)

const app = "remember"

var (
	log      = logging.MustGetLogger(app)
	RMBFILE  = os.Getenv("HOME") + "/.remember"
	remember *Remember
)

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
