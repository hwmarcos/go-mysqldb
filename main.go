package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mysqldb/config"
)

func main() {

	dbData := config.GetDataSource()

	db, err := sql.Open("mysql", ""+dbData[1]+":"+dbData[2]+"@tcp("+dbData[0]+":3306)/"+dbData[3]+"")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(db)

}
