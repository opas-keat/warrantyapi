package main

import (
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
	// zerolog.New(os.Stdout).With().Caller().Timestamp().Logger()
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
	// app.Use(logger.New(zerolog.New(os.Stdout).With().Caller().Timestamp().Logger()))
	app.Get("/metrics", monitor.New(monitor.Config{
		APIOnly: true,
	}))
	app.Use(recover.New())
	app.Use(requestid.New())

	//setup repository
	logRepository := repository.NewLogRepositoryImpl(database)
	dealerRepository := repository.NewDealerRepositoryImpl(database)
	productRepository := repository.NewProductRepositoryImpl(database)
	warrantyRepository := repository.NewWarrantyRepositoryImpl(database)
	configRepository := repository.NewConfigRepositoryImpl(database)
	userRepository := repository.NewUserRepositoryImpl(database)
	authenRepository := repository.NewAuthenRepositoryImpl(database)
	promotionRepository := repository.NewPromotionRepositoryImpl(database)

	//setup service
	dealerService := service.NewDealerServiceImpl(&dealerRepository, &logRepository)
	productService := service.NewProductServiceImpl(&productRepository, &logRepository)
	warrantyService := service.NewWarrantyServiceImpl(&warrantyRepository, &productRepository, &configRepository, &promotionRepository, &logRepository)
	configService := service.NewConfigServiceImpl(&configRepository, &logRepository)
	userService := service.NewUserServiceImpl(&userRepository, &logRepository)
	authenService := service.NewAuthenServiceImpl(&authenRepository)
	promotionService := service.NewPromotionServiceImpl(&promotionRepository, &logRepository)

	//setup controller
	dealerController := controller.NewDealerController(&dealerService, config)
	productController := controller.NewProductController(&productService, config)
	warrantyController := controller.NewWarrantyController(&warrantyService, config)
	fileController := controller.NewFileController(config)
	notificationController := controller.NewNotificationController(&warrantyService, config)
	configController := controller.NewConfigController(&configService, config)
	userController := controller.NewUserController(&userService, config)
	authenController := controller.NewAuthenController(&authenService, config)
	promotionController := controller.NewPromotionController(&promotionService, config)

	//setup routing
	app.Get("/", controller.Hello)
	app.Get("/healthz", controller.Hello)

	dealerController.Route(app)
	productController.Route(app)
	warrantyController.Route(app)
	fileController.Route(app)
	notificationController.Route(app)
	configController.Route(app)
	userController.Route(app)
	authenController.Route(app)
	promotionController.Route(app)
	app.All("*", controller.NotFound)

	// bytes := make([]byte, 32) //generate a random 32 byte key for AES-256
	// if _, err := rand.Read(bytes); err != nil {
	// 	panic(err.Error())
	// }
	//update
	// Close any connections on interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		_ = app.Shutdown()
	}()

	//start app
	err := app.Listen(config.Get("SERVER.HOST") + ":" + config.Get("SERVER.PORT"))
	exception.PanicLogging(err)

}
