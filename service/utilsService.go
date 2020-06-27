package service

import (
	"crypto/sha256"
	"fmt"
)

type Hash []byte

func (hash *Hash) String() string {
	return fmt.Sprintf("%x", *hash)
}

func GetHash(data string) []byte {
	hash := sha256.Sum256([]byte(data))
	return hash[:]
}
