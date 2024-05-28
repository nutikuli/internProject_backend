package datasource

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2/log"
)

type DbConfig struct {
	ServerPort int
	DbName     string
	DbUser     string
	DbPort     int
	DbPassword string
	DbHost     string
}

func GetDbConfig() DbConfig {

	serverPort := os.Getenv("SERVER_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")

	if serverPort == "" || dbName == "" || dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" {
		log.Info(serverPort, dbName, dbUser, dbPassword, dbHost, dbPort)
		panic("Missing required environment variables")
	}

	serverPortInt, err := strconv.Atoi(serverPort)
	if err != nil {
		panic("Invalid SERVER_PORT value")
	}
	dbPortInt, err := strconv.Atoi(dbPort)
	if err != nil {
		panic("Invalid DB_PORT value")
	}

	return DbConfig{
		ServerPort: serverPortInt,
		DbName:     dbName,
		DbUser:     dbUser,
		DbPort:     dbPortInt,
		DbPassword: dbPassword,
		DbHost:     dbHost,
	}

}
