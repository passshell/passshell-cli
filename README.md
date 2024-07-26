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
## Security

Passwords are encrypted using AES encryption before being stored. Ensure you keep your encryption key safe to maintain security.

## License

This project is licensed under the MIT License.

## Demo

![Image](.data/demo.gif)