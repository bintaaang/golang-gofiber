package encripyt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	//"fmt"
	"io"
)

// GenerateKeyBase64 menghasilkan key 32-byte (256 bit) dan mengembalikannya dalam Base64.
func GenerateKeyBase64() (string, error) {
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}

// EncryptToBase64 mengenkripsi plaintext menggunakan base64Key (harus 32-byte setelah decode).
// Output Base64 berisi: [nonce(12)][ciphertext][tag(16)]
func EncryptToBase64(plaintext string, base64Key string) (string, error) {
	// decode key
	key, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		return "", err
	}
	if len(key) != 32 {
		return "", errors.New("key harus 32 byte (Base64 dari 256-bit key)")
	}

	// buat cipher dan AEAD
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// generate nonce 12 byte
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Seal menghasilkan ciphertext||tag
	ct := aead.Seal(nil, nonce, []byte(plaintext), nil)

	// gabungkan nonce + ct(tag sudah termasuk di ct)
	out := append(nonce, ct...)

	// encode ke Base64
	return base64.StdEncoding.EncodeToString(out), nil
}

// DecryptFromBase64 mendekripsi base64Cipher dengan base64Key.
// Mengharapkan format input: [nonce(12)][ciphertext][tag(16)]
func DecryptFromBase64(base64Cipher string, base64Key string) (string, error) {
	// decode key
	key, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		return "", err
	}
	if len(key) != 32 {
		return "", errors.New("key harus 32 byte (Base64 dari 256-bit key)")
	}

	// decode cipher blob
	all, err := base64.StdEncoding.DecodeString(base64Cipher)
	if err != nil {
		return "", err
	}
	if len(all) < 12+16 {
		return "", errors.New("ciphertext tidak valid (terlalu pendek)")
	}

	nonce := all[:12]
	ct := all[12:] // sudah termasuk tag di akhir

	// buat cipher dan AEAD
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Open akan memverifikasi tag dan mengembalikan plaintext
	plaintext, err := aead.Open(nil, nonce, ct, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}