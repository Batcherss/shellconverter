package goshellconv

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/chacha20"
	"io"
)

func EncryptAES(data []byte, passphrase string) ([]byte, error) {
	key := []byte(passphrase)
	for len(key) < 16 {
		key = append(key, 0)
	}
	block, err := aes.NewCipher(key[:16])
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	encrypted := make([]byte, len(data))
	stream.XORKeyStream(encrypted, data)

	return append(iv, encrypted...), nil
}

func EncryptChaCha20(data []byte, passphrase string) ([]byte, error) {
	key := make([]byte, 32)
	copy(key, []byte(passphrase))
	if len(passphrase) < 32 {
		for i := len(passphrase); i < 32; i++ {
			key[i] = 0
		}
	}
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, err
	}
	encrypted := make([]byte, len(data))
	cipher.XORKeyStream(encrypted, data)
	return append(nonce, encrypted...), nil
}

func EncryptXOR(data []byte, passphrase string) ([]byte, error) {
	if len(passphrase) == 0 {
		return nil, errors.New("empty passphrase for XOR encryption")
	}
	key := []byte(passphrase)
	encrypted := make([]byte, len(data))

	for i := 0; i < len(data); i++ {
		encrypted[i] = data[i] ^ key[i%len(key)]
	}

	return encrypted, nil
}

func EncryptBase64(data []byte, _ string) ([]byte, error) {
	encoded := base64.StdEncoding.EncodeToString(data)
	return []byte(encoded), nil
}
