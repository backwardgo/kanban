package main

import (
	"log"
	"os"

	"github.com/backwardgo/kanban/httpd"
)

const usage = `
Usage:

	kanban [command]

Available commands are:

  httpd - starts an http server that provides access to our UI and API
`

func main() {
	var command string

	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	switch command {

	case "httpd":
		setPrefix(command)
		httpd.Run()

	default:
		log.Println(usage)
		os.Exit(1)

	}
}

func setPrefix(command string) {
	log.SetPrefix("[kanban::" + command + "] ")
}
