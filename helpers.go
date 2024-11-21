package stakego

import (
	"os"
	"strconv"
	"time"
)

// GetEnv - Get a string value from an environment variable
func GetEnv(env string, def string) string {
	v, exists := os.LookupEnv(env)
	if exists {
		return v
	}
	return def
}

// GetEnvInt - Get an int value from an environment variable
func GetEnvInt(env string, def int) int {
	v, exists := os.LookupEnv(env)
	if exists {
		v, err := strconv.Atoi(v)
		if err == nil {
			return v
		}
	}
	return def
}

// DatesEqual - compares just the date portion of a time.Time
func DatesEqual(t1 time.Time, t2 time.Time) bool {
	truncToDate := 24 * time.Hour
	return (t1.Truncate(truncToDate) == t2.Truncate(truncToDate))
}
