package encripyt

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const (
	SaltSize   = 16     // 128 bit
	KeySize    = 32     // 256 bit
	Iterations = 100000 // jumlah iterasi
)

// HashPassword membuat hash password dengan PBKDF2
func HashPassword(password string) (string, error) {
	// Generate salt random
	salt := make([]byte, SaltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// Derive key dengan PBKDF2
	key := pbkdf2.Key([]byte(password), salt, Iterations, KeySize, sha256.New)

	// Simpan salt:key dalam format base64
	saltB64 := base64.StdEncoding.EncodeToString(salt)
	keyB64 := base64.StdEncoding.EncodeToString(key)

	return fmt.Sprintf("%s:%s", saltB64, keyB64), nil
}

// VerifyPassword memverifikasi password dengan hash yang tersimpan
func VerifyPassword(password, storedHash string) (bool, error) {
	parts := strings.Split(storedHash, ":")
	if len(parts) != 2 {
		return false, fmt.Errorf("stored hash tidak valid (format harus 'salt:key')")
	}

	// Decode base64
	salt, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, err
	}

	key, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, err
	}

	// Derive ulang
	keyToCheck := pbkdf2.Key([]byte(password), salt, Iterations, KeySize, sha256.New)

	// Bandingkan secara aman (constant time)
	if subtle.ConstantTimeCompare(key, keyToCheck) == 1 {
		return true, nil
	}
	return false, nil
}
