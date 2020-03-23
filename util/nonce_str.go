package util

import "github.com/lingchengzeng/rand"

func NonceStr() string {
	return string(rand.NewHex())
}
