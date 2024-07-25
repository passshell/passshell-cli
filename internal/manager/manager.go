package manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"passshell/internal/crypto"

	"golang.org/x/term"
)

type PasswordManager struct {
	filename  string
	key       []byte
	passwords map[string]map[string]string
}

func New(filename string, key []byte) (*PasswordManager, error) {
	pm := &PasswordManager{
		filename:  filename,
		key:       key,
		passwords: make(map[string]map[string]string),
	}
	err := pm.load()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	return pm, nil
}

func GetSecureInput(prompt string) (string, error) {
	fmt.Print(prompt)
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		return "", err
	}
	return string(bytePassword), nil
}

func (pm *PasswordManager) AddPassword(service, username string) error {
	password, err := GetSecureInput("Enter password: ")
	if err != nil {
		return fmt.Errorf("error reading password: %v", err)
	}
	if _, ok := pm.passwords[service]; !ok {
		pm.passwords[service] = make(map[string]string)
	}
	pm.passwords[service][username] = password
	return pm.save()
}

func (pm *PasswordManager) GetPassword(service string) (map[string]string, error) {
	if users, ok := pm.passwords[service]; ok {
		return users, nil
	}
	return nil, fmt.Errorf("no passwords found for service: %s", service)
}

func (pm *PasswordManager) DeletePassword(service string) error {
	if _, ok := pm.passwords[service]; ok {
		delete(pm.passwords, service)
		return pm.save()
	}
	return errors.New("password not found")
}

func (pm *PasswordManager) ListServices() ([]string, error) {
	services := make([]string, 0, len(pm.passwords))
	for service := range pm.passwords {
		services = append(services, service)
	}
	return services, nil
}

func (pm *PasswordManager) save() error {
	data, err := json.Marshal(pm.passwords)
	if err != nil {
		return err
	}
	encryptedData, err := crypto.Encrypt(pm.key, data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(pm.filename, encryptedData, 0600)
}

func (pm *PasswordManager) load() error {
	data, err := ioutil.ReadFile(pm.filename)
	if err != nil {
		return err
	}
	decryptedData, err := crypto.Decrypt(pm.key, data)
	if err != nil {
		return err
	}
	return json.Unmarshal(decryptedData, &pm.passwords)
}
