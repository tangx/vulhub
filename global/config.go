package global

import (
	"database/sql"
	"os"
)

var (
	MYSQL_USER string = os.Getenv("MYSQL_USER")
	MYSQL_PASS string = os.Getenv("MYSQL_PASS")
	MYSQL_PORT string = os.Getenv("MYSQL_PORT")
	MYSQL_HOST string = os.Getenv("MYSQL_HOST")
	MysqlAgent *sql.DB
)

func init() {
	if MYSQL_HOST == "" {
		MYSQL_HOST = "mysql"
	}
	if MYSQL_PORT == "" {
		MYSQL_PORT = "3306"
	}
	if MYSQL_USER == "" {
		MYSQL_USER = "root"
	}
	if MYSQL_PASS == "" {
		MYSQL_PASS = "root123"
	}
}
