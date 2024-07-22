package main

import (
	"fmt"
	"log"

	"passshell/internal/crypto"
	"passshell/internal/manager"
	"passshell/pkg/cli"
)

func main() {
	// Запит майстер-пароля від користувача
	fmt.Print("Enter the master password: ")
	var masterPassword string
	fmt.Scanln(&masterPassword)

	// Генерація або завантаження ключа
	key, err := crypto.GetOrCreateKey(masterPassword)
	if err != nil {
		log.Fatalf("Error when receiving the key: %v", err)
	}

	// Ініціалізація менеджера паролів
	passwordManager, err := manager.New("passwords.json", key)
	if err != nil {
		log.Fatalf("PassShell failed to initialize: %v", err)
	}

	// Створення та запуск CLI
	c := cli.New(passwordManager)
	if err := c.Run(); err != nil {
		log.Fatalf("An error in the CLI: %v", err)
	}
}
