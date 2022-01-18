package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	engine := html.New("./views", ".html")
	// Return Fiber configuration.
	return fiber.Config{
		Views:       engine,
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}
