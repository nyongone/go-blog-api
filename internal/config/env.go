package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	DBUser				string
	DBPass				string
	DBHost				string
	DBPort				string
	DBSchema			string
	AppHost				string
	AppPort				string
	AppCors				string
	JWTSecret			string
	JWTATHour			string
	JWTRTHour			string
}

var EnvVar *Environment = new(Environment)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed loading environment variables: %v", err)
	}

	EnvVar.DBUser 		= os.Getenv("DB_USER")
	EnvVar.DBPass 		= os.Getenv("DB_PASS")
	EnvVar.DBHost 		= os.Getenv("DB_HOST")
	EnvVar.DBPort 		= os.Getenv("DB_PORT")
	EnvVar.DBSchema 	= os.Getenv("DB_SCHEMA")
	EnvVar.AppHost 		= os.Getenv("FIBER_APP_HOST")
	EnvVar.AppPort 		= os.Getenv("FIBER_APP_PORT")
	EnvVar.AppCors 		= os.Getenv("FIBER_APP_CORS_ORIGIN")
	EnvVar.JWTSecret	= os.Getenv("JWT_SECRET_KEY")
	EnvVar.JWTATHour	= os.Getenv("JWT_ACCESS_TOKEN_EXPIRY_HOUR")
	EnvVar.JWTRTHour	= os.Getenv("JWT_REFRESH_TOKEN_EXPIRY_HOUR")
}