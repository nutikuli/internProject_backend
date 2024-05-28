package datasource

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
)

func DbConnection() *sqlx.DB {
	config := GetDbConfig()

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)

	Db, err := sqlx.Open("mysql", connString)

	if err != nil {
		log.Error("**** Error creating connection pool: " + err.Error())
		panic(err.Error())
	}

	ctx := context.Background()
	err = Db.PingContext(ctx)
	if err != nil {
		fmt.Println("Catching ERR")
		log.Fatal(err.Error())
	}

	log.Debug("Connected!\n")

	return Db

}
