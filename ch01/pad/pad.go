// Package pad provides a one-time pad for encryption and decryption.
package pad

import (
	"math"
	"math/rand"
)

// A Key represents an encryption key.
type Key []byte

// A KeyPair represents a pair of encryption keys, typically a one-time pad and
// data encrypted with the pad.
type KeyPair struct {
	first  Key
	second Key
}

// randomKey creates a random one-time pad to act as a key.
func randomKey(length int) (result Key) {
	for i := 0; i < length; i++ {
		result = append(result, byte(rand.Intn(math.MaxUint8+1)))
	}
	return result
}

// xorKeys does an exclusive-or over two keys.
func xorKeys(first, second Key) (result Key) {
	for i, b := range first {
		result = append(result, b^second[i])
	}
	return result
}

// Encrypt encrypts original returning a key pair consisting of a random
// one-time pad and the encrypted data.
func Encrypt(original string) (result KeyPair) {
	result.first = randomKey(len(original))
	result.second = xorKeys(result.first, []byte(original))
	return result
}

// Decrypt decrypts a key pair consisting of a one-time pad and encrypted data.
func Decrypt(p KeyPair) string {
	result := xorKeys(p.first, p.second)
	return string(result)
}
