package service

import "context"

type passwordCrypter struct {
	ctx context.Context
}

type PasswordCrypter interface {
	Crypt() error
}

func NewPasswordCrypter(ctx context.Context) PasswordCrypter {
	return &passwordCrypter{
		ctx: ctx,
	}
}

func (pc *passwordCrypter) Crypt() error {
	return nil
}
