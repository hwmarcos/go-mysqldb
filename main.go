package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	dbData := GetDataSource()

	db, err := sql.Open("mysql", ""+dbData[1]+":"+dbData[2]+"@tcp("+dbData[0]+":3306)/"+dbData[3]+"")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(db)

}

func GetDataSource() []string {

	// load .env file
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
