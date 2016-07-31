package main

import (
	"log"
	"os"

	"github.com/backwardgo/kanban/db"
	"github.com/backwardgo/kanban/env"
	"github.com/backwardgo/kanban/httpd"
)

const usage = `Usage:

  kanban [command]

Available commands are:

  db:migrate:up
  db:migrate:down
  db:migrate:redo
  httpd

Use kanban help [command] for more information about a command.
`

func main() {
	var command string

	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	switch command {

	case "db:migrate:up":
		setPrefix(command)
		db.MigrateUp(env.DATABASE_URL)

	case "db:migrate:down":
		setPrefix(command)
		db.MigrateDown(env.DATABASE_URL)

	case "db:migrate:redo":
		setPrefix(command)
		db.MigrateRedo(env.DATABASE_URL)

	case "httpd":
		setPrefix(command)
		httpd.Run()

	default:
		showHelp()

	}
}

func showHelp() {
	log.SetFlags(0)

	var subcommand string

	if len(os.Args) > 2 {
		subcommand = os.Args[2]
	}

	if subcommand == "" {
		log.Println(usage)
		return
	}

	log.Printf(`Usage:

  kanban %s

`, subcommand)

	switch subcommand {

	case "db:migrate:up":
		log.Println(`apply all new database migrations`)

	case "db:migrate:down":
		log.Println(`undo the last database migration`)

	case "db:migrate:redo":
		log.Println(`db:migrate:down followed by db:migrate:up (useful for testing new migrations)`)

	case "httpd":
		log.Println(`starts an http server that provides access to our UI and API`)

	case "":
		log.Println(`requires a subcommand`)

	default:
		log.Printf(`%q is an unkown (or undocumented) command`, subcommand)

	}
}

func setPrefix(command string) {
	log.SetPrefix("[kanban::" + command + "] ")
}
