package coreengine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
)

// Validate validates a struct using the validator library.
func Validate(i interface{}) error {
	validate := validator.New()
	return validate.Struct(i)
}

// SaveToFile saves a struct to a file.
func SaveToFile(i interface{}, filePath string) error {
	data, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

// LoadFromFile loads a struct from a file.
func LoadFromFile(filePath string, i interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, i)
}

// UUID generates a random UUID.
func UUID() (uuid.UUID, error) {
	return uuid.NewV4()
}

// Now returns the current time in seconds since the epoch.
func Now() int64 {
	return time.Now().Unix()
}

// SameUUID checks if two UUIDs are the same.
func SameUUID(a, b uuid.UUID) bool {
	return a == b
}

// GetPath returns the path to a file.
func GetPath(base, relativePath string) (string, error) {
	return filepath.Abs(filepath.Join(base, relativePath))
}

// SplitPath splits a file path into its base and relative path components.
func SplitPath(path string) (string, string, error) {
	base := filepath.Dir(path)
	relativePath := filepath.Base(path)
	return base, relativePath, nil
}

// IsFile checks if a path is a file.
func IsFile(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return stat.Mode().IsRegular(), nil
}

// IsDir checks if a path is a directory.
func IsDir(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return stat.Mode().IsDir(), nil
}

// GetParentDir returns the parent directory of a file path.
func GetParentDir(path string) (string, error) {
	return filepath.Dir(path), nil
}

// GetFileName returns the file name of a file path.
func GetFileName(path string) (string, error) {
	return filepath.Base(path), nil
}

// GetDirName returns the directory name of a file path.
func GetDirName(path string) (string, error) {
	return filepath.Dir(path), nil
}

// HasExtension checks if a file path has a specific extension.
func HasExtension(path, ext string) (bool, error) {
	fileExt := filepath.Ext(path)
	return fileExt == ext, nil
}

// IsEmptyString checks if a string is empty.
func IsEmptyString(s string) bool {
	return s == ""
}

// SanitizeString removes non-alphanumeric characters from a string.
func SanitizeString(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' ||
			r >= 'A' && r <= 'Z' ||
			r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, s)
}

// SanitizeUsername removes non-alphanumeric characters and underscores from a string.
func SanitizeUsername(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' ||
			r >= 'A' && r <= 'Z' ||
			r >= '0' && r <= '9' ||
			r == '_' {
			return r
		}
		return -1
	}, s)
}

// IsEmailValid checks if an email is valid.
func IsEmailValid(email string) bool {
	return strings.Contains(email, "@")
}

// SentryCaptureException captures an exception and sends it to Sentry.
func SentryCaptureException(ex error) {
	if sentry.Init(sentry.ClientOptions{
		EnableTraces: true,
	}) != nil {
		// Panic and exit if initialization fails
		panic("unable to start Sentry")
	}
	sentry.CaptureException(ex)
}

// PanicIfError panics if the given error is not nil.
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

// Errorf returns a formatted error message.
func Errorf(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}