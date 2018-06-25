package helpers

// GenerateRandomBytes returns securely generated random bytes.
import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateToken generates a base64-encoded securely generated random string.
func GenerateToken(length uint8) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}
