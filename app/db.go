package app

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnectionDb() *sql.DB {

	db, err := sql.Open("mysql", "root:rifkiganteng@tcp(localhost:3306)/restful?parseTime=true")
	if err != nil {
		fmt.Println(err, "error disini")
		panic(err)
	}

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(50 * time.Minute)
	return db
}
