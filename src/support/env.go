package support

import "os"

/* Get environment variable with fallback */
func Env(key string, fallback string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return fallback
}
