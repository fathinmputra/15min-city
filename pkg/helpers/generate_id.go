package helpers

import (
	"time"
)

func GenerateId() string {
	currentTime := time.Now().UTC().Format("20060102150405")
	return currentTime
}
