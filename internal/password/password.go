package password

import (
	"fmt"
	"math/rand/v2"
)

type Generator interface {
	Generate() (string, error)
}

type Password struct {
	length        int
	useDigits     bool
	useLowercase  bool
	useUppercase  bool
	usedPasswords map[string]bool
}

func New(length int, useDigits, useLowercase, useUppercase bool) *Password {
	return &Password{
		length:        length,
		useDigits:     useDigits,
		useLowercase:  useLowercase,
		useUppercase:  useUppercase,
		usedPasswords: make(map[string]bool),
	}
}

var (
	digits    = []rune("0123456789")
	lowercase = []rune("abcdefghijklmnopqrstuvwxyz")
	uppercase = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func (p *Password) Generate() (string, error) {
	if p.length <= 0 {
		return "", fmt.Errorf("lenth must be positive")
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

	result := make([]rune, p.length)
	for i := 0; i < p.length; i++ {
		result[i] = charPool[rand.IntN(len(charPool))]
	}

	if p.useDigits {
		result[rand.IntN(p.length)] = digits[rand.IntN(len(digits))]
	}
	if p.useLowercase {
		result[rand.IntN(p.length)] = lowercase[rand.IntN(len(lowercase))]
	}
	if p.useUppercase {
		result[rand.IntN(p.length)] = uppercase[rand.IntN(len(uppercase))]
	}

	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	password := string(result)

	if p.usedPasswords[password] {
		return p.Generate()
	}

	p.usedPasswords[password] = true
	return password, nil
}
