// Package env provides functionality for loading and managing environment variables from .env files.
//
// The package supports loading all environment variables from a file or retrieving specific
// variables individually. It handles common .env file features including comments, empty lines,
// and quoted values (both single and double quotes).
//
// Example usage:
//
//	// Load all environment variables from file
//	err := env.LoadEnv(".env")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Get a specific environment variable
//	value, err := env.GetEnv("DB_HOST", ".env")
//	if err != nil {
//		log.Fatal(err)
//	}
package env

import (
	"bufio"
	"os"
	"strings"
)

// LoadEnv reads environment variables from a file and sets them in the environment.
// Each line in the file should be in KEY=VALUE format. The function supports:
//
// - Comments (lines starting with #)
// - Empty lines
// - Quoted values (both single and double quotes)
// - Basic KEY=VALUE format
//
// Lines that don't conform to the KEY=VALUE format are silently skipped.
//
// Example .env file content:
//
//	# Database settings
//	DB_HOST=localhost
//	DB_PORT=5432
//	APP_NAME="My Application"
//	API_KEY='secret-key'
//
// Returns an error if the file cannot be opened or read.
func LoadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		value = strings.Trim(value, `"'`)

		os.Setenv(key, value)
	}

	return scanner.Err()
}

// GetEnv retrieves the value of a specific environment variable from the given file.
// It follows the same parsing rules as LoadEnv but only returns the value for the
// specified key.
//
// The function will:
// - Skip comment lines (starting with #)
// - Skip empty lines
// - Remove surrounding quotes (both single and double) from values
// - Return the first matching value if the key appears multiple times
//
// If the key is not found, it returns an empty string and nil error.
// Returns an error only if the file cannot be opened or read.
func GetEnv(key string, filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		if strings.TrimSpace(parts[0]) == key {
			value := strings.TrimSpace(parts[1])
			return strings.Trim(value, `"'`), nil
		}
	}

	return "", scanner.Err()
}
