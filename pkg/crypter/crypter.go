package crypter

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type crypter struct {
	key []byte
}

func New(key string) *crypter {
	return &crypter{
		key: []byte(key),
	}
}

func (c *crypter) Encrypt(text string) (string, error) {
	textByte := []byte(text)
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", errors.Join(ErrBlockCreate, err)
	}

	padding := aes.BlockSize - len(textByte)%aes.BlockSize
	padText := append(textByte, bytes.Repeat([]byte{byte(padding)}, padding)...)

	ciphertext := make([]byte, aes.BlockSize+len(padText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", errors.Join(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], padText)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (c *crypter) Decrypt(ciptext string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(ciptext)
	if err != nil {
		return "", errors.Join(ErrInvalidBase64, err)
	}
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", errors.Join(ErrBlockCreate, err)
	}
	if len(ciphertext) < aes.BlockSize {
		return "", errors.Join(ErrTextToShort, err)
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	padding := int(ciphertext[len(ciphertext)-1])
	return string(ciphertext[:len(ciphertext)-padding]), nil
}
