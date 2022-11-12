package conf

import (
	"log"
	"os"
	"strconv"
)

const (
	hostKey       = "BV_HOST"
	portKey       = "BV_PORT"
	dbHostKey     = "BV_DB_HOST"
	dbPortKey     = "BV_DB_PORT"
	dbNameKey     = "BV_DB_NAME"
	dbUserKey     = "BV_DB_USER"
	dbPasswordKey = "BV_DB_PASSWORD"
)

type Config struct {
	Host       string
	Port       string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
}

func NewConfig() Config {
	host, ok := os.LookupEnv(hostKey)
	if !ok || host == "" {
		logAndPanic(hostKey)
	}

	port, ok := os.LookupEnv(portKey)
	if !ok || port == "" {
		if _, err := strconv.Atoi(port); err != nil {
			logAndPanic(portKey)
		}
	}

	dbHost, ok := os.LookupEnv(dbHostKey)
	if !ok || dbHost == "" {
		logAndPanic(dbHostKey)
	}

	dbPort, ok := os.LookupEnv(dbPortKey)
	if !ok || dbPort == "" {
		if _, err := strconv.Atoi(dbPort); err != nil {
			logAndPanic(dbPortKey)
		}
	}

	dbName, ok := os.LookupEnv(dbNameKey)
	if !ok || dbName == "" {
		logAndPanic(dbNameKey)
	}

	dbUser, ok := os.LookupEnv(dbUserKey)
	if !ok || dbUser == "" {
		logAndPanic(dbUserKey)
	}

	dbPassword, ok := os.LookupEnv(dbPasswordKey)
	if !ok || dbPassword == "" {
		logAndPanic(dbPasswordKey)
	}

	return Config{
		Host:       host,
		Port:       port,
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbName:     dbName,
		DbUser:     dbUser,
		DbPassword: dbPassword,
	}
}

func logAndPanic(envVar string) {
	log.Println("ENV variable not set or value not valid: ", envVar)
	panic(envVar)
}
