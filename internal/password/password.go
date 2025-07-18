package password

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type Generator interface {
	Generate() (string, error)
}

type Password struct {
	length       int
	useDigits    bool
	useLowercase bool
	useUppercase bool
	uniqueChars  bool
}

func New(length int, useDigits, useLowercase, useUppercase bool, uniqueChars bool) *Password {
	return &Password{
		length:       length,
		useDigits:    useDigits,
		useLowercase: useLowercase,
		useUppercase: useUppercase,
		uniqueChars:  uniqueChars,
	}
}

var (
	digits    = []rune("0123456789")
	lowercase = []rune("abcdefghijklmnopqrstuvwxyz")
	uppercase = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func (p *Password) Generate() (string, error) {
	if p.length < 4 || p.length > 40 {
		return "", fmt.Errorf("length must be between 4 and 40")
	}
	if !p.useDigits && !p.useLowercase && !p.useUppercase {
		return "", fmt.Errorf("select at least one character type")
	}

	var charPool []rune
	if p.useDigits {
		charPool = append(charPool, digits...)
	}
	if p.useLowercase {
		charPool = append(charPool, lowercase...)
	}
	if p.useUppercase {
		charPool = append(charPool, uppercase...)
	}

	if p.uniqueChars && p.length > len(charPool) {
		return "", fmt.Errorf("can't generate password: length is greater than the number of unique characters")
	}

	if p.uniqueChars {

		for attempt := 0; attempt < 10; attempt++ {
			shuffled := make([]rune, len(charPool))
			copy(shuffled, charPool)
			cryptoShuffle(shuffled)
			candidate := shuffled[:p.length]

			hasDigit := !p.useDigits
			hasLower := !p.useLowercase
			hasUpper := !p.useUppercase
			for _, c := range candidate {
				if !hasDigit && containsRune(digits, c) {
					hasDigit = true
				}
				if !hasLower && containsRune(lowercase, c) {
					hasLower = true
				}
				if !hasUpper && containsRune(uppercase, c) {
					hasUpper = true
				}
			}
			if hasDigit && hasLower && hasUpper {
				return string(candidate), nil
			}
		}
		return "", fmt.Errorf("failed to generate password with unique characters and all selected types")
	}

	result := make([]rune, p.length)

	types := [][]rune{}
	if p.useDigits {
		types = append(types, digits)
	}
	if p.useLowercase {
		types = append(types, lowercase)
	}
	if p.useUppercase {
		types = append(types, uppercase)
	}
	for i, t := range types {
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(t))))
		result[i] = t[idx.Int64()]
	}
	start := len(types)
	for i := start; i < p.length; i++ {
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charPool))))
		result[i] = charPool[idx.Int64()]
	}
	cryptoShuffle(result)
	return string(result), nil
}

func cryptoShuffle(runes []rune) {
	n := len(runes)
	for i := n - 1; i > 0; i-- {
		jBig, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		j := int(jBig.Int64())
		runes[i], runes[j] = runes[j], runes[i]
	}
}

func containsRune(pool []rune, r rune) bool {
	for _, c := range pool {
		if c == r {
			return true
		}
	}
	return false
}
