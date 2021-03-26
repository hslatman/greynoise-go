package responses

import "fmt"

type Error struct {
	Code    int
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s (%d)", e.Message, e.Code)
}
