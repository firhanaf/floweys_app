package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
)

var JWT_SECRET string = ""

type AppConfig struct {
	DBUsername string
	DBPass     string
	DBHost     string
	DBPORT     int
	DBName     string
	JWTSecret  string
}

func InitConfig() *AppConfig { return ReadENV() }

func ReadENV() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("DB_USERNAME"); found {
		app.DBUsername = val
		isRead = false
	}

	if val, found := os.LookupEnv("DB_PASS"); found {
		app.DBPass = val
		isRead = false
	}

	if val, found := os.LookupEnv("DB_HOST"); found {
		app.DBHost = val
		isRead = false
	}

	if val, found := os.LookupEnv("DB_PORT"); found {
		conVal, _ := strconv.Atoi(val)
		app.DBPORT = conVal
		isRead = false
	}

	if val, found := os.LookupEnv("DB_NAME"); found {
		app.DBName = val
		isRead = false
	}

	if val, found := os.LookupEnv("JWT_SECRET"); found {
		app.JWTSecret = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Fatal error config file: %s \n", err)
		}
		app.DBName = viper.GetString("DB_NAME")
		app.DBHost = viper.GetString("DB_HOST")
		app.DBPORT = viper.GetInt("DB_PORT")
		app.DBUsername = viper.GetString("DB_USERNAME")
		app.DBPass = viper.GetString("DB_PASS")
		app.JWTSecret = viper.GetString("JWT_SECRET")
	}
	JWT_SECRET = app.JWTSecret
	return &app
}
