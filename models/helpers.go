package models

import (
	"fmt"
	"strings"
)

const (
	minPasswordLen = 6
)

var (
	errorf  = fmt.Errorf
	sprintf = fmt.Sprintf

	toLower   = strings.ToLower
	trimSpace = strings.TrimSpace
)
