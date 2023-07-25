package util

import (
	"github.com/Nikkoz/mp.gateway/internal/domain/store/types/name"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomUInt(min, max int64) uint64 {
	return uint64(min + rand.Int63n(max-min+1))
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomStoreName() name.Name {
	return name.Name(RandomString(6))
}

func RandomStoreToken() string {
	return RandomString(10)
}

func RandomStoreAuthToken() string {
	return RandomString(12)
}

func RandomStoreClientId() string {
	return RandomString(16)
}

func RandomStoreClientSecret() string {
	return RandomString(20)
}
