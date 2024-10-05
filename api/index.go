package handler

import (
	"net/http"
	"os"

	"github.com/ahargunyllib/hackathon-fiber-starter/internal/infra/database"
	"github.com/ahargunyllib/hackathon-fiber-starter/internal/infra/env"
	"github.com/ahargunyllib/hackathon-fiber-starter/internal/infra/server"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

// Handler is the main entry point of the application. Think of it like the main() method
func Handler(w http.ResponseWriter, r *http.Request) {
	// This is needed to set the proper request path in `*fiber.Ctx`
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

// building the fiber application
func handler() http.HandlerFunc {
	env.GetEnv()

	server := server.NewHttpServer()
	psqldb := database.NewPgsqlConn()

	database.Migrate(psqldb, os.Args)
	database.Seeder(psqldb, os.Args)

	server.MountMiddlewares()
	server.MountRoutes(psqldb)

	app := server.GetApp()

	return adaptor.FiberApp(app)
}
