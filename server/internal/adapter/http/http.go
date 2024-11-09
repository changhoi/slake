package http

import (
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recovery "github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

func NewAdapter(l *zap.Logger) *fiber.App {

	errLogger := l.Named("http")
	app := fiber.New(fiber.Config{
		JSONDecoder: json.Unmarshal,
		JSONEncoder: json.Marshal,
	})

	app.Use(helmet.New())
	app.Use(recovery.New(recovery.Config{
		StackTraceHandler: func(_ *fiber.Ctx, e interface{}) {
			errLogger.Error("panic", zap.Any("error", e))
		},
	}))
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{}))

	return app
}
