package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func generateKey(secret string) []byte {
	hash := sha256.New()
	hash.Write([]byte(secret))
	return hash.Sum(nil)
}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func EncryptPass(password string) (string, error) {
	key := generateKey(os.Getenv("SECRET_KEY"))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("error creating AES Block cipher: %w", err)
	}

	plaintext := []byte(password)

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("error generating random IV: %w", err)
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)

	result := append(iv, ciphertext...)
	return encodeBase64(result), nil
}

func Decrypt(encrypted string) (string, error) {
	key := generateKey(os.Getenv("SECRET_KEY"))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("error creating AES Block cipher: %w", err)
	}

	data, err := decodeBase64(encrypted)
	if err != nil {
		return "", fmt.Errorf("error decoding base64 data: %w", err)
	}

	if len(data) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := data[:aes.BlockSize]
	ciphertext := data[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}
