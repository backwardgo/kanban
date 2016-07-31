package db

import "fmt"

type direction string

const (
	Ascending  direction = "ASC"
	Descending direction = "DESC"
)

type OrderBy struct {
	tableName string
	fieldName string
	direction direction
}

func (ob OrderBy) String() string {
	return fmt.Sprintf("%v.%v %v", ob.tableName, ob.fieldName, ob.direction)
}
