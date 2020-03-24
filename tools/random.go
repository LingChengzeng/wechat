package tools

import (
	"math/rand"
	"time"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func Random(n int) []byte {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}

		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}

		cache >>= letterIdxBits
		remain--
	}

	return b
}

// RandomString accepts a number(10 for example) and returns a random
// string using simple but fairly safe random algorithm
func RandomString(n int) string {
	return string(Random(n))
}

// 生成随机数 [min, max)
func RandInt64(min, max int64) int64 {
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min >= max || max == 0 {
		return 0
	} else if min == 0 {
		return rander.Int63n(max)
	}
	return rander.Int63n(max-min) + min
}
