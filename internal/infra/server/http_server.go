package server

import (
	"github.com/ahargunyllib/hackathon-fiber-starter/internal/infra/env"
	"github.com/ahargunyllib/hackathon-fiber-starter/internal/middlewares"
	"github.com/ahargunyllib/hackathon-fiber-starter/pkg/bcrypt"
	errorhandler "github.com/ahargunyllib/hackathon-fiber-starter/pkg/helpers/http/error_handler"
	"github.com/ahargunyllib/hackathon-fiber-starter/pkg/helpers/http/response"
	"github.com/ahargunyllib/hackathon-fiber-starter/pkg/log"
	timePkg "github.com/ahargunyllib/hackathon-fiber-starter/pkg/time"
	"github.com/ahargunyllib/hackathon-fiber-starter/pkg/uuid"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type HttpServer interface {
	Start(part string)
	MountMiddlewares()
	MountRoutes(db *sqlx.DB)
	GetApp() *fiber.App
}

type httpServer struct {
	app *fiber.App
}

func NewHttpServer() HttpServer {
	config := fiber.Config{
		Prefork:       env.AppEnv.AppEnv == "production",
		CaseSensitive: true,
		AppName:       "Hackathon Fiber Starter",
		ServerHeader:  "Hackathon Fiber Starter",
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
		ErrorHandler:  errorhandler.ErrorHandler,
	}

	app := fiber.New(config)

	return &httpServer{
		app: app,
	}
}

func (s *httpServer) GetApp() *fiber.App {
	return s.app
}

func (s *httpServer) Start(port string) {
	if port[0] != ':' {
		port = ":" + port
	}

	err := s.app.Listen(port)

	if err != nil {
		log.Fatal(log.LogInfo{
			"error": err.Error(),
		}, "[SERVER][Start] failed to start server")
	}
}

func (s *httpServer) MountMiddlewares() {
	s.app.Use(middlewares.LoggerConfig())
	s.app.Use(middlewares.Helmet())
	s.app.Use(middlewares.Compress())
	s.app.Use(middlewares.Cors())
	if env.AppEnv.AppEnv != "development" {
		s.app.Use(middlewares.ApiKey())
	}
	s.app.Use(middlewares.RecoverConfig())
}

func (s *httpServer) MountRoutes(db *sqlx.DB) {
	_ = bcrypt.Bcrypt
	_ = timePkg.Time
	_ = uuid.UUID

	_ = middlewares.NewMiddleware()

	s.app.Get("/", func(c *fiber.Ctx) error {
		return response.SendResponse(c, fiber.StatusOK, "hai manies😘")
	})

	api := s.app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return response.SendResponse(c, fiber.StatusOK, "hai manies😘")
	})

	s.app.Use(func(c *fiber.Ctx) error {
		return c.SendFile("./web/not-found.html")
	})
}
