package cryptoutils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetHash(url string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(url))
	return hex.EncodeToString(hash.Sum(nil))
}
