package database

import (
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/ahargunyllib/hackathon-fiber-starter/internal/infra/env"
	"github.com/ahargunyllib/hackathon-fiber-starter/pkg/log"
	"github.com/jmoiron/sqlx"
)

func NewPgsqlConn() *sqlx.DB {
	driverName := env.AppEnv.DBConnection
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ", env.AppEnv.DBHost, env.AppEnv.DBPort, env.AppEnv.DBUser, env.AppEnv.DBPass, env.AppEnv.DBName)

	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		log.Panic(log.LogInfo{
			"error": err.Error(),
		}, "[DB][NewPgsqlConn] failed to connect to database")
		panic(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
