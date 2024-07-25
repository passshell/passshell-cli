package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"passshell/internal/manager"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

type CLI struct {
	manager *manager.PasswordManager
}

func New(manager *manager.PasswordManager) *CLI {
	return &CLI{manager: manager}
}

func (c *CLI) printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  add                            	 - Add a new entry")
	fmt.Println("  get <service>         		 - Get entry")
	fmt.Println("  rm  <service>                   	 - Remove password")
	fmt.Println("  ls                              	 - Show list of entryes")
	fmt.Println("  help                            	 - Show this help")
	fmt.Println("  clear                               	 - Clear passshell window")
	fmt.Println("  exit                            	 - Exit the passshell")
}

func Welcome() {
	myFigure := figure.NewFigure("PassShell", "", true)
	myFigure.Print()
	fmt.Println("Welcome to PassShell!")
	fmt.Println("Type 'help' to view the available commands.")
}

func Goodbye() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	myFigure := figure.NewFigure("Goodbye!", "", true)
	myFigure.Print()
}

func (c *CLI) Run() error {
	Welcome()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			Goodbye()
			fmt.Println("Thank you for using PassShell. Goodbye!")
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
			if len(parts) != 1 {
				fmt.Println("Usage: add")
				continue
			}
			fmt.Print("Enter service: ")
			service, _ := reader.ReadString('\n')
			service = strings.TrimSpace(service)

			fmt.Print("Enter username: ")
			username, _ := reader.ReadString('\n')
			username = strings.TrimSpace(username)

			err := c.manager.AddPassword(service, username)
			if err != nil {
				fmt.Printf("Error adding password: %v\n", err)
			} else {
				fmt.Println("Entry successfully added!")
			}

		case "get":
			if len(parts) != 2 {
				fmt.Println("Usage: get <service>")
				continue
			}
			passwords, err := c.manager.GetPassword(parts[1])
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Printf("Entryes for service '%s':\n", parts[1])
				for username, password := range passwords {
					fmt.Printf("Username: %s\nPassword: %s\n", username, password)
				}
			}

		case "rm":
			if len(parts) != 2 {
				fmt.Println("Usage: rm <service>")
				continue
			}
			fmt.Printf("Are you sure you want to delete all entryes for service '%s'? (y/n): ", parts[1])

			var confirmation string
			_, err := fmt.Scanln(&confirmation)
			if err != nil {
				fmt.Printf("Error reading confirmation: %v\n", err)
				continue
			}
			if strings.ToLower(strings.TrimSpace(confirmation)) != "y" {
				fmt.Println("Deletion cancelled.")
				continue
			}
			err = c.manager.DeletePassword(parts[1])
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Service and all associated entryes successfully deleted.")
			}

		case "ls":
			services, err := c.manager.ListServices()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else if len(services) > 0 {
				fmt.Println("Available entryes:")
				for _, service := range services {
					fmt.Println(service)
				}
			} else {
				fmt.Println("You have not available entryes yet.")
			}

		case "clear":
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			Welcome()

		default:
			fmt.Println("Unknown command. Use 'help' to see the available commands.")
		}
	}
}
