package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const (
	RandomSize = 3
	Random     = 123
)

func SchoolCodeGeneretor(name string) string {
	// Return the first 6 characters of the password
	return fmt.Sprintf("%s%d", strings.ToUpper(name), RandomNumber(RandomSize))
}

func RandomNumber(length int) int64 {
	max := int64(1)
	for i := 0; i < length; i++ {
		max *= 10
	}

	b, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return Random
	}

	return b.Int64()
}
