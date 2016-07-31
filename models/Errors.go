package models

import (
	"encoding/json"
	"log"
)

func NewErrors() Errors {
	return make(Errors)
}

type Errors map[string]string

func (e Errors) Blank() bool {
	return len(e) == 0
}

func (e Errors) Error() string {
	return e.String()
}

func (e Errors) Present() bool {
	return !e.Blank()
}

func (e Errors) String() string {
	if e.Blank() {
		return `null`
	}

	data, err := json.Marshal(e)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
