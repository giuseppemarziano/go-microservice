package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/palantir/stacktrace"
	"golang.org/x/crypto/argon2"
)

type passwordHasher struct {
	ctx       context.Context
	cryptCost int
}

type PasswordHasher interface {
	Hash(password string) (string, error)
}

func NewPasswordHasher(ctx context.Context, cryptCost int) PasswordHasher {
	return &passwordHasher{
		ctx:       ctx,
		cryptCost: cryptCost,
	}
}

func (pc *passwordHasher) Hash(password string) (string, error) {
	if password == "" {
		return "", stacktrace.Propagate(
			errors.New("password cannot be empty"), // TODO add domain error
			"error on hashing",
		)
	}

	salt := make([]byte, pc.cryptCost)
	if _, err := rand.Read(salt); err != nil {
		return "", stacktrace.Propagate(
			err,
			"error on generating random salt for password hashing",
		)
	}

	hashedPassword := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedPassword := base64.RawStdEncoding.EncodeToString(hashedPassword)
	return "$argon2id$" + encodedSalt + "$" + encodedPassword, nil
}
