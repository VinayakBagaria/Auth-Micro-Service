package db

import (
	"log"
	"os"
	"strconv"
)

type Config interface {
	Dsn() string
	DbName() string
}

type config struct {
	dbUser string
	dbPass string
	dbHost string
	dbPort int
	dbName string
	dsn    string
}

func NewConfig() Config {
	var cfg config
	cfg.dbUser = os.Getenv("DATABASE_USER")
	cfg.dbPass = os.Getenv("DATABASE_PASS")
	cfg.dbHost = os.Getenv("DATABASE_HOST")
	var err error
	cfg.dbPort, err = strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Fatalln("Error on load env var: ", err.Error())
	}
	cfg.dbName = os.Getenv("DATABASE_NAME")
	// cfg.dsn = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authMechanism=SCRAM-SHA-1&authSource=%s", cfg.dbUser, cfg.dbPass, cfg.dbHost, cfg.dbPort, "", cfg.dbName)
	cfg.dsn = "mongodb://user:password@localhost:27017/?authSource=microservices"
	return &cfg
}

func (c *config) Dsn() string {
	return c.dsn
}

func (c *config) DbName() string {
	return c.dbName
}
