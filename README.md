# Go env

A lightweight Go library for loading and managing environment variables from .env files.

## Features

- Load environment variables from .env files
- Retrieve specific environment variables from .env files
- Support for comments and empty lines
- Support for quoted values (both single and double quotes)
- Simple and easy to use API

## Installation

```bash
go get github.com/cesp99/go-env
```

## Usage

### Loading Environment Variables

```go
package main

import (
    "fmt"
    "github.com/cesp99/go-env"
)

func main() {
    // Load all environment variables from .env file
    err := env.LoadEnv(".env")
    if err != nil {
        fmt.Printf("Error loading .env file: %v\n", err)
        return
    }
}
```

### Getting Specific Environment Variables

```go
package main

import (
    "fmt"
    "github.com/cesp99/go-env"
)

func main() {
    // Get a specific environment variable from .env file
    value, err := env.GetEnv("DB_HOST", ".env")
    if err != nil {
        fmt.Printf("Error reading environment variable: %v\n", err)
        return
    }
    fmt.Printf("DB_HOST: %s\n", value)
}
```

### Example .env File

```env
# Database settings
DB_HOST=localhost
DB_PORT=5432

# Application settings
APP_NAME="My Application"
API_KEY='secret-key'
```

## Features

### Comments and Empty Lines
- Lines starting with `#` are treated as comments
- Empty lines are ignored

### Quoted Values
- Supports both single and double quoted values
- Quotes are automatically trimmed from the value

### Error Handling
- Returns appropriate errors for file operations
- Skips malformed lines without failing

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the GPL 3.0 License - see the [LICENSE](./LICENSE.txt) file for details.
