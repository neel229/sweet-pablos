package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphabets = "abcdefghijklmnopqrstuvwxyz"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random number
// between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomInt32 generates a random number
// between min and max
func RandomInt32(min, max int64) int32 {
	return int32(min + rand.Int63n(max-min+1))
}

// RandomString generates a random string
// of length of n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomBool returns true if
// rand.Intn returns 1;else false
func RandomBool() bool {
	return rand.Intn(2) == 1
}
