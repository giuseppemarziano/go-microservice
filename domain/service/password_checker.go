package service

import "context"

type passwordChecker struct {
	ctx context.Context
}

type PasswordChecker interface {
	Check() error
}

func NewPasswordChecker(ctx context.Context) PasswordChecker {
	return &passwordChecker{
		ctx: ctx,
	}
}

func (pc *passwordChecker) Check() error {
	return nil
}
