package error

import "fmt"

type PasswordTooShort struct {
	MinLength int
}

func (p PasswordTooShort) Error() string {
	return fmt.Sprintf("password must be at least %d characters long", p.MinLength)
}
