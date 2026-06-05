# đź” gopwd

![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)

A lightning-fast, cryptographically secure password generator CLI written in Go.

## Features
- Cryptographically secure random generation via `crypto/rand`
- Flexible customization (length, character sets)
- Zero external dependencies
- Cross-platform support

## Installation
```bash
go install github.com/username/gopwd@latest
```

## Usage
```bash
# Generate a 16-character password (default)
$ gopwd
K8#p@L9v!mQ2x$Z1

# Generate a 32-character password without special characters
$ gopwd -l 32 -s=false
aB3dE5fG7hI9jK1lM3nO5pQ7rS9tU1vW
```
