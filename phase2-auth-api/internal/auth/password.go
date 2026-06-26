package auth

import "golang.org/x/crypto/bcrypt"

// TODO(phase2-auth-password): Implement password hashing with bcrypt.
//
// Task:
//   1. HashPassword(plain string) (string, error) — use bcrypt.GenerateFromPassword with cost 12.
//   2. CheckPassword(hash, plain string) error — use bcrypt.CompareHashAndPassword.
//   3. Never log or return plaintext passwords.
//
// Learn: https://pkg.go.dev/golang.org/x/crypto/bcrypt
// Allowed dep: golang.org/x/crypto/bcrypt (only non-stdlib dependency in this project).

func HashPassword(plain string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CheckPassword(hash, plain string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	if err != nil {
		return err
	}

	return nil
}
