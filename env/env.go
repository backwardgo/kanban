package env

import "os"

var (
	DATABASE_URL = os.Getenv("DATABASE_URL")
	KANBAN_HOME  = os.Getenv("KANBAN_HOME")
	PORT         = os.Getenv("PORT")
)
