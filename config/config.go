package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort int
	DBName string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}
	isRead := true
	if val, found := os.LookupEnv("DBUSER"); found {
		app.DBUser = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		app.DBPass = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DBHost = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.DBPort = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DBName = val
		isRead = false
	}

	if isRead {
		err := godotenv.Load("local.env")
		if err != nil {
			fmt.Println("Error saat baca env", err.Error())
			return nil
		}

		app.DBUser = os.Getenv("DBUSER")
		app.DBPass = os.Getenv("DBPASS")
		app.DBHost = os.Getenv("DBHOST")
		readData := os.Getenv("DBPORT")
		app.DBPort, err = strconv.Atoi(readData)
		if err != nil {
			fmt.Println("Error saat convert DBPORT", err.Error())
			return nil
		}
		app.DBName = os.Getenv("DBNAME")
	}

	return &app
}
