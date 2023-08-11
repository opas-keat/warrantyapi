package configuration

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		AppName:      "PPSW-WARRANTY API",
		ServerHeader: "PPSW-WARRANTY",
		BodyLimit:    10 * 1024 * 1024,
		JSONEncoder:  sonic.Marshal,
		JSONDecoder:  sonic.Unmarshal,
		// ErrorHandler:      exception.ErrorHandler,
		EnablePrintRoutes: false,
	}
}
