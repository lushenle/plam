package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomEmail gen generates a random
func RandomEmail() string {
	return RandomString(6) + "@" + RandomString(6) + ".com"
}

// RandomFloat32 generates a random float32 between min and max
func RandomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
