package common

import (
	"crypto/sha256"
	"fmt"

	"github.com/google/uuid"
)

func NewUUIDString() string {
	uuidRaw := uuid.New()
	return uuidRaw.String()
}

func Encrypt(plainText string) (crypted string) {
	asBytes := sha256.Sum256([]byte(plainText))
	crypted = fmt.Sprintf("%x", asBytes)
	return
}
