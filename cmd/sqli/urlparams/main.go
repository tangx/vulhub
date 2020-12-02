package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/tangx/vulhub/modules/mysqldb"
)

func main() {
	GinServer()
}

func init() {
	db := mysqldb.Conn()

	createDatabase(db)
	createTable(db)
}

func createDatabase(db *sql.DB) {
	_, err := db.Exec(`CREATE DATABASE IF NOT EXISTS vulhub ;`)
	if err != nil {
		logrus.Errorf("create database failed: %s\n", err.Error())
		panic(err)
	}

	_, _ = db.Exec(`use vulhub;`)
}

func createTable(db *sql.DB) {

	sql := `CREATE TABLE IF NOT EXISTS users  (
			id int(0) NOT NULL AUTO_INCREMENT,
			name varchar(255) NULL,
			password varchar(255) NULL,
			PRIMARY KEY (id)
		);
	`

	logrus.Debugf("%s\n", sql)

	if _, err := db.Exec(string(sql)); err != nil {
		// logrus.Errorf("create table failed: %s\n", err.Error())
		logrus.Errorf("init database failed: %s\n", err.Error())
		return
	}
	logrus.Infoln("init database succeeded")

	_, _ = db.Exec(`INSERT INTO users(name, password) VALUES ('admin', 'admin') ;`)
	_, _ = db.Exec(`INSERT INTO users(name, password) VALUES ('zhangsan', '123456') ;`)
	_, _ = db.Exec(`INSERT INTO users(name, password) VALUES ('lisi', '223456') ;`)
}
