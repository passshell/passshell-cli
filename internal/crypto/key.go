package crypto

import (
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/scrypt"
)

const keyFile = "key.encrypted"

func GetOrCreateKey(masterPassword string) ([]byte, error) {
	if _, err := os.Stat(keyFile); os.IsNotExist(err) {
		return createAndSaveKey(masterPassword)
	}
	return loadAndDecryptKey(masterPassword)
}

func createAndSaveKey(masterPassword string) ([]byte, error) {
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, err
	}

	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, err
	}

	derivedKey, err := deriveKey([]byte(masterPassword), salt)
	if err != nil {
		return nil, err
	}

	encryptedKey, err := Encrypt(derivedKey, key)
	if err != nil {
		return nil, err
	}

	data := append(salt, encryptedKey...)
	if err := ioutil.WriteFile(keyFile, data, 0600); err != nil {
		return nil, err
	}

	return key, nil
}

func loadAndDecryptKey(masterPassword string) ([]byte, error) {
	data, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}

	salt, encryptedKey := data[:16], data[16:]

	derivedKey, err := deriveKey([]byte(masterPassword), salt)
	if err != nil {
		return nil, err
	}

	return Decrypt(derivedKey, encryptedKey)
}

func deriveKey(password, salt []byte) ([]byte, error) {
	return scrypt.Key(password, salt, 32768, 8, 1, 32)
}
