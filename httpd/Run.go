package httpd

import (
	"log"
	"net/http"

	"github.com/backwardgo/kanban/env"
)

func Run() {
	port := ":" + env.PORT
	log.Printf("listen and serve on port %s", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("http.ListenAndServe failed %q", err)
	}
}
