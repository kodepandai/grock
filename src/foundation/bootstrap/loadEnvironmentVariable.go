package bootstrap

import (
	"grock/src/foundation"

	"github.com/joho/godotenv"
)

type LoadEnvironmentVariable struct{}

func (l LoadEnvironmentVariable) Bootstrap(app *foundation.Application) {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}
}
