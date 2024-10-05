package main

import (
	"os"

	"github.com/ahargunyllib/hackathon-fiber-starter/internal/infra/database"
	"github.com/ahargunyllib/hackathon-fiber-starter/internal/infra/env"
	"github.com/ahargunyllib/hackathon-fiber-starter/internal/infra/server"
	"github.com/ahargunyllib/hackathon-fiber-starter/pkg/log"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal(log.LogInfo{"error": err}, "Error loading .env file")
	}

	env.GetEnv()

	server := server.NewHttpServer()
	psqldb := database.NewPgsqlConn()

	database.Migrate(psqldb, os.Args)
	database.Seeder(psqldb, os.Args)

	server.MountMiddlewares()
	server.MountRoutes(psqldb)
	server.Start(env.AppEnv.AppPort)
}
