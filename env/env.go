package env

import "os"

var (
	PORT = os.Getenv("PORT")
)
