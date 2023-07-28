package ext

import (
	"math/rand"
	"regexp"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const colorRegex = `/^#([0-9A-F]{3}){1,2}$/i`

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
func IsColorHex(n string) bool {
	matched, _ := regexp.MatchString(colorRegex, n)
	if matched {
		return matched
	}
	return false
}
