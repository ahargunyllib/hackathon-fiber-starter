package env

import (
	"os"
	"time"

	"github.com/ahargunyllib/hackathon-fiber-starter/pkg/log"
)

type Env struct {
	AppEnv            string        `mapstructure:"APP_ENV"`
	AppPort           string        `mapstructure:"APP_PORT"`
	ApiKey            string        `mapstructure:"API_KEY"`
	DBHost            string        `mapstructure:"DB_HOST"`
	DBPort            string        `mapstructure:"DB_PORT"`
	DBUser            string        `mapstructure:"DB_USER"`
	DBPass            string        `mapstructure:"DB_PASS"`
	DBName            string        `mapstructure:"DB_NAME"`
	JwtSecretKey   []byte        `mapstructure:"JWT_SECRET_KEY"`
	JwtExpTime time.Duration `mapstructure:"JWT_EXP_TIME"`
}

var AppEnv *Env

func GetEnv() {
	env := &Env{}

	env.AppPort = os.Getenv("APP_PORT")
	env.AppEnv = os.Getenv("APP_ENV")
	env.ApiKey = os.Getenv("API_KEY")
	env.DBHost = os.Getenv("DB_HOST")
	env.DBPort = os.Getenv("DB_PORT")
	env.DBUser = os.Getenv("DB_USER")
	env.DBPass = os.Getenv("DB_PASS")
	env.DBName = os.Getenv("DB_NAME")
	env.JwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	dur, err := time.ParseDuration(os.Getenv("JWT_EXP_TIME"))
	if err != nil {
		log.Fatal(log.LogInfo{"error": err.Error()}, "Fail to parse JWT_EXP_TIME")
	}
	env.JwtExpTime = dur

	switch env.AppEnv {
	case "development":
		log.Info(nil, "Application is running on development mode")
	case "production":
		log.Info(nil, "Application is running on production mode")
	case "staging":
		log.Info(nil, "Application is running on staging mode")
	}

	AppEnv = env
}
