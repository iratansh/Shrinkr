package services

import (
	"crypto/rand"
	"math/big"
)

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// Shortener generates short codes for long URLs.
// Strategy: random base62 codes of configurable length, with collision retry.
type Shortener struct {
	CodeLength int
}

// NewShortener returns a Shortener that emits random base62 codes.
func NewShortener(codeBytes int) *Shortener {
	if codeBytes <= 0 {
		codeBytes = 6
	}
	return &Shortener{CodeLength: codeBytes}
}

// Generate returns a new random short code.
func (s *Shortener) Generate() (string, error) {
	code := make([]byte, s.CodeLength)
	max := big.NewInt(int64(len(alphabet)))

	for i := range code {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		code[i] = alphabet[n.Int64()]
	}

	return string(code), nil
}
