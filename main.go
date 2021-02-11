package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mysqldb/config"
)

func main() {

	dbData := config.GetDataSource()

	dataSourceName := "" + dbData.User + ":" + dbData.Pass + "@tcp(" + dbData.Host + ":3306)/" + dbData.DbName + ""

	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(db)

}
