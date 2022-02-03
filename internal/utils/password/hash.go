package password

import (
	"crypto/sha256"
	"encoding/base64"
)

func HashPassword(password string) string{
	hasher := sha256.New()
	hasher.Write([]byte(password))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}