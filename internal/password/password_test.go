package password

import (
	"testing"
)

func TestPasswordLength(t *testing.T) {
	gen := New(12, true, true, true, true)
	pw, err := gen.Generate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len([]rune(pw)) != 12 {
		t.Errorf("expected length 12, got %d", len([]rune(pw)))
	}
}

func TestPasswordTypes(t *testing.T) {
	gen := New(10, true, true, true, true)
	pw, err := gen.Generate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	hasDigit, hasLower, hasUpper := false, false, false
	for _, c := range pw {
		if containsRune(digits, c) {
			hasDigit = true
		}
		if containsRune(lowercase, c) {
			hasLower = true
		}
		if containsRune(uppercase, c) {
			hasUpper = true
		}
	}
	if !hasDigit || !hasLower || !hasUpper {
		t.Errorf("password should contain all types: got %q", pw)
	}
}

func TestUniqueChars(t *testing.T) {
	gen := New(20, true, true, true, true)
	pw, err := gen.Generate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	seen := map[rune]bool{}
	for _, c := range pw {
		if seen[c] {
			t.Errorf("character %q is repeated in password %q", c, pw)
		}
		seen[c] = true
	}
}

func TestTooLongUnique(t *testing.T) {
	gen := New(11, true, false, false, true)
	_, err := gen.Generate()
	if err == nil {
		t.Error("expected error for too long unique password, got nil")
	}
}

func TestNoTypesSelected(t *testing.T) {
	gen := New(8, false, false, false, true)
	_, err := gen.Generate()
	if err == nil {
		t.Error("expected error for no character types selected, got nil")
	}
}
