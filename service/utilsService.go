package service

import (
	"crypto/sha256"
	"flag"
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

func Testing() bool {
	return flag.Lookup("test.v") != nil
}
