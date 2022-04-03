package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/fatih/color"
)

// Creates a small startup delay
// For example helpful for waiting for DB container
// (e.g. in the case of Docker Compose)
func StartupDelay(duration time.Duration) {
	message := fmt.Sprintf("Starting in %d seconds...", int(duration.Seconds()))
	color.Set(color.FgYellow, color.Bold)
	log.Println(message)
	color.Unset()
	time.Sleep(duration)
}

const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Generates a random string that is used as the unique part of a shortlink.
func GenerateRandomShortString(length int) (string, error) {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes), nil
}

// Checks whether or not the passed string is a valid
// URL.
//
// Needed for instance for validating long and short links
// and for checking whether they're okay to be created/deleted
// or okay to be further proccessed.
func IsStringValidURL(str string) bool {
	// Try to parse string as raw URL,
	// if that's not possible, the string is not
	// a valid URL
	_, err := url.ParseRequestURI(str)
	if err != nil {
		// Not a valid URL
		return false
	}

	// Try to parse string as URL
	// If anything mandatory is missing in the parsed URL,
	// the string is not a valid URL
	u, err := url.Parse(str)
	if err != nil || u.Scheme == "" || u.Host == "" {
		// Not a valid URL
		return false
	}

	// The link seems to be a valid URL
	return true
}
