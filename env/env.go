package env

import "os"

var (
	DATABASE_URL  = os.Getenv("DATABASE_URL")
	KANBAN_HOME   = os.Getenv("KANBAN_HOME")
	PASSWORD_SALT = os.Getenv("PASSWORD_SALT")
	PORT          = os.Getenv("PORT")
)
