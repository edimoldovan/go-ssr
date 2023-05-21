package config

import (
	"os"
)

func IsDevelopment() bool {
	return os.Getenv("GO_WEB_ENV") == "development"
}
