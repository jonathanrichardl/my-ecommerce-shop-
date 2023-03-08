package util

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"regexp"

	"github.com/google/uuid"
)

func NewUUID() string {
	return uuid.New().String()
}

func ValidateEmail(input string) error {
	emailIsValid, err := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, input)
	if err != nil {
		return err
	}

	if emailIsValid {
		return nil
	}

	return errors.New("Invalid email")
}

func EncryptPassword(password string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
}
