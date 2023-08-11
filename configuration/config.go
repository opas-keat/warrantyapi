package configuration

import (
	"os"
	"warrantyapi/exception"

	"github.com/joho/godotenv"
)

const Secret = "omsoft"

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	exception.PanicLogging(err)
	return &configImpl{}
}
