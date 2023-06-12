package main

import (
	"database/sql"

	"github.com/celsopires1999/matchreport/configs"
	_ "github.com/lib/pq"
)

func main() {
	configs := configs.LoadConfig(".")
	conn, err := sql.Open("postgres", configs.DBConn)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
}
