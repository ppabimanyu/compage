package password

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	password := "mySecret123!"
	hashed, err := GeneratePassword(password)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if hashed == password {
		t.Errorf("hashed password should not be the same as the original password")
	}
	// Check if the hash is valid
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		t.Errorf("generated hash does not match the password: %v", err)
	}
}

func TestComparePassword(t *testing.T) {
	password := "anotherSecret!"
	hashed, err := GeneratePassword(password)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	// Should not return error for correct password
	if err := ComparePassword(hashed, password); err != nil {
		t.Errorf("expected passwords to match, got error: %v", err)
	}
	// Should return error for wrong password
	wrongPassword := "wrongPassword"
	if err := ComparePassword(hashed, wrongPassword); err == nil {
		t.Errorf("expected error for wrong password, got nil")
	}
}
