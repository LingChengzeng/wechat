package util

import "github.com/LingChengzeng/rand"

func NonceStr() string {
	return string(rand.NewHex())
}
