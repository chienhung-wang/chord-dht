package util

import (
	"crypto/sha1"
	"math/big"
)

func Sha1_addr(addr string) *big.Int {
	var b big.Int

	hash_bytes := sha1.Sum([]byte(addr))

	b.SetBytes(hash_bytes[:])

	return &b
}
