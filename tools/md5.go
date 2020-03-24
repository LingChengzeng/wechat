// md5
package tools

import (
	"crypto/md5"
	"encoding/hex"
)

// md5
func Md5hash(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
