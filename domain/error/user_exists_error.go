package error

import "fmt"

type EmailAlreadyExists struct {
	Email string
}

func (e EmailAlreadyExists) Error() string {
	return fmt.Sprintf("email already exists: %s", e.Email)
}
