package main

import (
	"crypto/rand"
	"os"
	"os/signal"
	"time"
	"warrantyapi/configuration"
	"warrantyapi/controller"
	"warrantyapi/exception"
	repository "warrantyapi/repository/impl"
	service "warrantyapi/service/impl"

	"github.com/rs/zerolog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	//setup configuration
	config := configuration.New()
	database := configuration.NewDatabase(config)
	configuration.NewDatabase(config)
	zerolog.TimeFieldFormat = time.RFC3339
	//setup fiber
	app := fiber.New(configuration.NewFiberConfiguration())

	//setup fiber middleware
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(cors.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You have requested too many in a single time-frame! Please wait another minute!",
			})
		},
	}))
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] | ${pid} | ${locals:requestid} | ${status} - ${latency} | ${method} | ${path}​\n",
		TimeFormat: "02-01-2006 15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))
	app.Get("/metrics", monitor.New(monitor.Config{
		APIOnly: true,
	}))
	app.Use(recover.New())
	app.Use(requestid.New())

	//setup repository
	logRepository := repository.NewLogRepositoryImpl(database)
	dealerRepository := repository.NewDealerRepositoryImpl(database)

	//setup service
	dealerService := service.NewDealerServiceImpl(&dealerRepository, &logRepository)

	//setup controller
	dealerController := controller.NewDealerController(&dealerService, config)

	//setup routing
	app.Get("/", controller.Hello)
	app.Get("/healthz", controller.Hello)
	dealerController.Route(app)
	app.All("*", controller.NotFound)

	bytes := make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	// Close any connections on interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		_ = app.Shutdown()
	}()

	//start app
	err := app.Listen(config.Get("SERVER.PORT"))
	exception.PanicLogging(err)

}
