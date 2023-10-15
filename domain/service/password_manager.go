package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/palantir/stacktrace"
	"go-microservice/domain/repositories"
	"golang.org/x/crypto/argon2"
	"strings"
)

type PasswordManager interface {
	Check(email string, password string) error
	Hash(password string) (string, error)
}

type passwordManager struct {
	ctx            context.Context
	cryptCost      int
	userRepository repositories.UserRepository
}

func (pc *passwordManager) Check(email string, password string) error {
	email = strings.ToLower(email)

	user, err := pc.userRepository.GetUserByEmail(pc.ctx, email)
	if err != nil {
		return stacktrace.Propagate(
			err,
			"error on retrieving user with email: %s",
			email,
		)
	}
	if user == nil {
		return stacktrace.Propagate(
			errors.New("user not found"),
			"error on finding user with email: %s",
			email,
		)
	}

	parts := strings.Split(user.Password, "$")
	if len(parts) != 3 {
		return stacktrace.Propagate(
			errors.New("invalid password format"),
			"error on validating stored password format for user: %s",
			email,
		)
	}
	salt, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return stacktrace.Propagate(
			err,
			"error on decoding salt for user: %s",
			email,
		)
	}
	storedHash, err := base64.RawStdEncoding.DecodeString(parts[2])
	if err != nil {
		return stacktrace.Propagate(
			err,
			"error on decoding hashed password for user: %s",
			email,
		)
	}

	hashedPassword := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	if !comparePasswords(hashedPassword, storedHash) {
		return stacktrace.Propagate(
			errors.New("password mismatch"),
			"error on password verification for user: %s",
			email,
		)
	}

	return nil
}

func comparePasswords(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	var result byte
	for i := 0; i < len(a); i++ {
		result |= a[i] ^ b[i]
	}
	return result == 0
}

func (pc *passwordManager) Hash(password string) (string, error) {
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
