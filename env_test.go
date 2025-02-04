package env

import (
	"os"
	"testing"
)

func createTempEnvFile(content string) (string, error) {
	tmpfile, err := os.CreateTemp("", "env-test-*.env")
	if err != nil {
		return "", err
	}
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		os.Remove(tmpfile.Name())
		return "", err
	}
	if err := tmpfile.Close(); err != nil {
		os.Remove(tmpfile.Name())
		return "", err
	}
	return tmpfile.Name(), nil
}

func TestLoadEnv(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantErr bool
	}{
		{
			name: "basic variables",
			content: `DB_HOST=localhost
DB_PORT=5432
`,
			wantErr: false,
		},
		{
			name: "with comments and empty lines",
			content: `# Database settings
DB_HOST=localhost

DB_PORT=5432
`,
			wantErr: false,
		},
		{
			name: "with quoted values",
			content: `MESSAGE="Hello, World!"
PATH='usr/local/bin'
`,
			wantErr: false,
		},
		{
			name: "malformed entries",
			content: `KEY1=value1
INVALID_LINE
KEY2=value2
`,
			wantErr: false,
		},
		{
			name:    "empty file",
			content: "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename, err := createTempEnvFile(tt.content)
			if err != nil {
				t.Fatalf("failed to create temp file: %v", err)
			}
			defer os.Remove(filename)

			err = LoadEnv(filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadEnv() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Verify environment variables were set correctly
			if tt.name == "basic variables" {
				if got := os.Getenv("DB_HOST"); got != "localhost" {
					t.Errorf("DB_HOST = %v, want localhost", got)
				}
				if got := os.Getenv("DB_PORT"); got != "5432" {
					t.Errorf("DB_PORT = %v, want 5432", got)
				}
			}
		})
	}

	// Test non-existent file
	if err := LoadEnv("non_existent_file.env"); err == nil {
		t.Error("LoadEnv() expected error for non-existent file")
	}
}

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name    string
		content string
		key     string
		wantVal string
		wantErr bool
	}{
		{
			name: "existing key",
			content: `TEST_KEY=test_value
`,
			key:     "TEST_KEY",
			wantVal: "test_value",
			wantErr: false,
		},
		{
			name: "non-existing key",
			content: `OTHER_KEY=value
`,
			key:     "TEST_KEY",
			wantVal: "",
			wantErr: false,
		},
		{
			name: "quoted value",
			content: `QUOTED_KEY="quoted value"
`,
			key:     "QUOTED_KEY",
			wantVal: "quoted value",
			wantErr: false,
		},
		{
			name: "with comments",
			content: `# Comment
TEST_KEY=test_value
`,
			key:     "TEST_KEY",
			wantVal: "test_value",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename, err := createTempEnvFile(tt.content)
			if err != nil {
				t.Fatalf("failed to create temp file: %v", err)
			}
			defer os.Remove(filename)

			gotVal, err := GetEnv(tt.key, filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotVal != tt.wantVal {
				t.Errorf("GetEnv() = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}

	// Test non-existent file
	_, err := GetEnv("ANY_KEY", "non_existent_file.env")
	if err == nil {
		t.Error("GetEnv() expected error for non-existent file")
	}
}
