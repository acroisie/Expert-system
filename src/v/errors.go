package v

import (
	"fmt"
)

type ErrorType string

const (
	CONTRADICTION ErrorType = "CONTRADICTION"
	FACT_NOT_FOUND ErrorType = "FACT_NOT_FOUND"
	SOLVING ErrorType = "SOLVING"
	MAX_DEPTH ErrorType = "MAX_DEPTH"
)

type Error struct {
	Type ErrorType
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}
