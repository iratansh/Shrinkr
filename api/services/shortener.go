package services

// Shortener generates short codes for long URLs.
// Strategy: random base62 codes of configurable length, with collision retry.
type Shortener struct {
	CodeBytes int
}

// NewShortener returns a Shortener that emits codes derived from codeBytes random bytes.
func NewShortener(codeBytes int) *Shortener {
	if codeBytes <= 0 {
		codeBytes = 6
	}
	return &Shortener{CodeBytes: codeBytes}
}

// Generate returns a new random short code.
func (s *Shortener) Generate() (string, error) {
	// TODO: read crypto/rand bytes and base62-encode them.
	return "", nil
}
