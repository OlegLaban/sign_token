package crypter

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

type Crypter interface {
	Encrypt(string) ([]byte, error)
	Decrypt(string) (string, error)
}

type crypter struct {
	key []byte
}

func New(key string) *crypter {
	return &crypter{
		key: []byte(key),
	}
}

func (c *crypter) Encrypt(text string) ([]byte, error) {
	textByte := []byte(text)
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, errors.Join(ErrBlockCreate, err)
	}

	padding := aes.BlockSize - len(textByte)%aes.BlockSize
	padText := append(textByte, bytes.Repeat([]byte{byte(padding)}, padding)...)

	ciphertext := make([]byte, aes.BlockSize+len(padText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, errors.Join(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], padText)

	return ciphertext, nil
}

func (c *crypter) Decrypt(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, errors.Join(ErrBlockCreate, err)
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.Join(ErrTextToShort, err)
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	padding := int(ciphertext[len(ciphertext)-1])
	return ciphertext[:len(ciphertext)-padding], nil
}
