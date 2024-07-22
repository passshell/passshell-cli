# PassShell - CLI Password Manager

A simple and efficient command-line password manager written in Go.

## Features

- **Add passwords**: Easily store new passwords.
- **Retrieve passwords**: Securely access your stored passwords.
- **List all stored services and usernames**: View a comprehensive list of all your entries.
- **Delete passwords**: Remove passwords from storage when they are no longer needed.
- **Help**: Display a list of available commands and their usage.

## Usage

### 1. Build the Project

Compile the program using the following command:

```sh
go build -o passshell cmd/main.go
```
### 2. Run the Program

Execute the compiled program with:

```sh
./passshell
```
### 3. Commands

    add <service> <username> <password>: Add a new password.
    get <service> <username>:            Retrieve a password.
    ls:                                  List all stored services and usernames.
    rm <service> <username>:             Delete a stored password.
    help:                                Display a list of available commands and their usage.
    exit:                                Exit the program.

## Security

Passwords are encrypted using AES encryption before being stored. Ensure you keep your encryption key safe to maintain security.

## License

This project is licensed under the MIT License.

## Example

```sh
$ ./password-manager
> add github myusername myGithubPassword123
Password added successfully
> add email myemail@example.com myEmailPassword456
Password added successfully
> list
github | myusername
email | myemail@example.com
> get github myusername
Password: myGithubPassword123
> delete github myusername
Password deleted successfully
> list
email | myemail@example.com
> help
Available commands:
  add <service> <username> <password>           - Add a new password
  get <service> <username>                      - Retrieve a password
  list                                          - List all stored services and usernames
  delete <service> <username>                   - Delete a stored password
  help                                          - Display this help message
  exit                                          - Exit the program
> exit

```