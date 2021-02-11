package config

import (
	"github.com/joho/godotenv"
	"os"
)

func GetDataSource() []string {

	err := godotenv.Load(".env")

	if err != nil {
		panic(err.Error())
	}

	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	dbname := os.Getenv("MYSQL_DBNAME")

	return []string{host, user, pass, dbname}
}
