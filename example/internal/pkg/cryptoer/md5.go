package cryptoer

import (
	"crypto/md5"
	"fmt"
)

func (m *Manager) ToMd5(msg string) string {
	return fmt.Sprintf("%x",
		md5.Sum(
			[]byte(msg),
		),
	)
}
