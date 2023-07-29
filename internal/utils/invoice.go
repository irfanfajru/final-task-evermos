package utils

import (
	"math/rand"
	"time"
)

func GenerateInvoiceCode() string {
	const charset = "0123456789"
	const length = 10
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
