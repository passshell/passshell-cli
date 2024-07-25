package main

import (
	"log"
	"os"
	"os/exec"

	"passshell/internal/crypto"
	"passshell/internal/manager"
	"passshell/pkg/cli"
)

func main() {
	masterPassword, err := manager.GetSecureInput("Enter the master password: ")
	if err != nil {
		log.Fatalf("Error reading password: %v", err)
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	key, err := crypto.GetOrCreateKey(masterPassword)
	if err != nil {
		log.Fatalf("Error when receiving the key: %v", err)
	}

	passwordManager, err := manager.New("passwords.json", key)
	if err != nil {
		log.Fatalf("PassShell failed to initialize: %v", err)
	}

	c := cli.New(passwordManager)
	if err := c.Run(); err != nil {
		log.Fatalf("An error in the CLI: %v", err)
	}
}
