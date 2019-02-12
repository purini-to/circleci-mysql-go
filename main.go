package main

import (
	"database/sql"
	"github.com/purini-to/circleci-mysql-go/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/circle_test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // test
	service.Query(db)
}
