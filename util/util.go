package util

import (
	"crypto/sha1"
	"math/big"
)

func Sha1_hash(addr string) *big.Int {
	var b big.Int

	hash_bytes := sha1.Sum([]byte(addr))

	b.SetBytes(hash_bytes[:])

	return &b
}

func Between(begin *big.Int, mid *big.Int, end *big.Int) bool {
	if end.Cmp(begin) > 0 {
		return mid.Cmp(begin) > 0 && mid.Cmp(end) <= 0
	} else {
		return mid.Cmp(begin) > 0 || mid.Cmp(end) <= 0
	}
}

func BetweenNoninclusive(begin *big.Int, mid *big.Int, end *big.Int) bool {
	if end.Cmp(begin) > 0 {
		return mid.Cmp(begin) > 0 && mid.Cmp(end) < 0
	} else {
		return mid.Cmp(begin) > 0 || mid.Cmp(end) < 0
	}
}
