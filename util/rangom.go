package util

import (
	"math/rand"
	"strings"
	"time"
)

const alpabet = "qwertyuiopasdfghjklzxcvbnm"

func init() {
	rand.Seed(time.Now().Unix())

}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alpabet)

	for i := 0; i < k; i++ {
		c := alpabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()

}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 10e5)
}

func RandomCurrency() string {
	currencies := []string{"RUB", "USD", "EUR"}
	return currencies[rand.Intn(len(currencies)-1)]
}
