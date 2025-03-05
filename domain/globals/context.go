package globals

import (
	"crypto/rand"
	"fmt"
)

func GenerateAPIKey() (string, error) {
	key := make([]byte, 12)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", key), nil
}
