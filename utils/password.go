package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Returns the bcrypt hash of the password.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// Check if the provided password is correct or not.
func CheckPassword(password string, hashedPAssword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPAssword), []byte(password))
}
