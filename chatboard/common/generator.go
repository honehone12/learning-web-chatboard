package common

import "github.com/google/uuid"

func NewUUIDString() string {
	uuidRaw := uuid.New()
	return uuidRaw.String()
}
