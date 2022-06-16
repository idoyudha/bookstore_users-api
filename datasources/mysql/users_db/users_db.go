package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/idoyudha/bookstore_utils-go/logger"
	"github.com/joho/godotenv"
)

const (
	mysqlUsersUsername = "MYSQL_USERS_USERNAME"
	mysqlUsersPassword = "MYSQL_USERS_PASSWORD"
	mysqlUsersHost     = "MYSQL_USERS_HOST"
	mysqlUsersSchema   = "MYSQL_USERS_SCHEMA"
)

var (
	Client *sql.DB
)

func init() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		os.Getenv(mysqlUsersUsername),
		os.Getenv(mysqlUsersPassword),
		os.Getenv(mysqlUsersHost),
		os.Getenv(mysqlUsersSchema),
	)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	mysql.SetLogger(logger.GetLogger())
	log.Println("database successfully configured!")
}
