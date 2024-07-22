package cli

import (
	"bufio"
	"fmt"
	"os"
	"passshell/internal/manager"
	"strings"
)

type CLI struct {
	manager *manager.PasswordManager
}

func New(manager *manager.PasswordManager) *CLI {
	return &CLI{manager: manager}
}

func (c *CLI) printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  add <service> <login> <password>	 - Add a new password")
	fmt.Println("  get <service> <login> 		 - Get password")
	fmt.Println("  rm  <service> <login>           	 - Remove password")
	fmt.Println("  ls                              	 - Show list of services")
	fmt.Println("  help                            	 - Show this help")
	fmt.Println("  exit                            	 - Exit the program")
}

func (c *CLI) Run() error {
	fmt.Println("Welcome to PassShell!")
	fmt.Println("Type 'help' to view the available commands.")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			return nil
		}

		parts := strings.Split(input, " ")
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "help":
			c.printHelp()

		case "add":
			if len(parts) != 4 {
				fmt.Println("Usage: add <service> <username> <password>")
				continue
			}
			err := c.manager.AddPassword(parts[1], parts[2], parts[3])
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Entry successfully added!")
			}

		case "get":
			if len(parts) != 3 {
				fmt.Println("Usage: get <service> <login>")
				continue
			}
			password, err := c.manager.GetPassword(parts[1], parts[2])
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Printf("Password: %s\n", password)
			}

		case "rm":
			if len(parts) != 3 {
				fmt.Println("Usage: rm <service> <login>")
				continue
			}
			err := c.manager.DeletePassword(parts[1], parts[2])
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Entry successfully deleted")
			}

		case "ls":
			services, err := c.manager.ListServices()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Available services:")
				for _, service := range services {
					fmt.Println(service)
				}
			}

		default:
			fmt.Println("Unknown command. Use 'help' to see the available commands.")
		}
	}
}
