package password

import "crypto/sha256"

func HashPassword(password string) string{
	hash := sha256.Sum256([]byte(password))
	return string(hash[:])
}