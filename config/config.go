package config

import (
	"os"
	"path"

	"guardian/common"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	currentWD, err := os.Getwd()
	common.PanicOnError(err)
	common.PanicOnError(godotenv.Load(path.Join(currentWD, ".env")))
}
