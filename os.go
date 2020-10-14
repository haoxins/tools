package tools

import "os"

// Getenv get env with a default value
func Getenv(key, fallback string) string {
	v := os.Getenv(key)

	if v == "" {
		v = fallback
	}

	return v
}
