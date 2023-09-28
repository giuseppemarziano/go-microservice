package error

import "fmt"

type UserNotFound struct {
	Email string
}

func (p UserNotFound) Error() string {
	return fmt.Sprintf("user not found with email: %s", p.Email)
}
