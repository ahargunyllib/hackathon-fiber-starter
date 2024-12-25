package main

import (
	"github.com/ahargunyllib/hackathon-fiber-starter/internal/infra/database"
	"github.com/ahargunyllib/hackathon-fiber-starter/internal/infra/env"
	"github.com/ahargunyllib/hackathon-fiber-starter/internal/infra/server"
)

// @title						Hackathon Fiber Starter API
// @version					1.0
// @description				This is Hackathon Fiber Starter API Documentation
// @host						localhost:8080
// @securityDefinitions.apiKey	UserAuth
// @in							header
// @name						Authorization
// @description				API Key for accessing protected user and admin endpoints. Type: Bearer TOKEN
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						x-api-key
// @description				API Key for accessing all endpoints. Type: Key TOKEN
func main() {
	env.GetEnv()

	server := server.NewHttpServer()
	psqlDB := database.NewPgsqlConn()
	defer psqlDB.Close()

	server.MountMiddlewares()
	server.MountRoutes(psqlDB)
	server.Start(env.AppEnv.AppPort)
}
