package service

import (
	"crypto/sha256"
	"encoding/hex"
)

// Sum256 sha256加密
func Sum256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
