package stakego

import (
  "os"
  "strconv"
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