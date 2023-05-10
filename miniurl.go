package miniurl

import (
	"crypto/md5"
	"encoding/hex"
)

func Hast(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
