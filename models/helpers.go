package models

import (
	"fmt"
	"strings"
)

var (
	errorf  = fmt.Errorf
	sprintf = fmt.Sprintf

	toLower   = strings.ToLower
	trimSpace = strings.TrimSpace
)
