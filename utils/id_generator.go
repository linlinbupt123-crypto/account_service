package utils

import (
	"github.com/google/uuid"
)

func GenerateTxID() string {
	return uuid.New().String()
}
