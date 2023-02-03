package ecrypto

import (
	"crypto/md5"
	"fmt"
)

func ToMd5(msg string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(msg)))
}
