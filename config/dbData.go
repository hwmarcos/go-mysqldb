package config

import (
	"github.com/joho/godotenv"
	"os"
)

type dbData struct {
	Host   string
	User   string
	Pass   string
	DbName string
}

func GetDataSource() dbData {

	err := godotenv.Load(".env")

	if err != nil {
		panic(err.Error())
	}

	return dbData{
		Host:   os.Getenv("MYSQL_HOST"),
		User:   os.Getenv("MYSQL_USER"),
		Pass:   os.Getenv("MYSQL_PASS"),
		DbName: os.Getenv("MYSQL_DBNAME"),
	}
}
