package mysqldb

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/tangx/vulhub/global"
)

func Open() *sql.DB {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/`, global.MYSQL_USER, global.MYSQL_PASS,
		global.MYSQL_HOST, global.MYSQL_PORT)

	logrus.Infof("dsn = %s\n", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超时的连接就close
	db.SetMaxOpenConns(100)                  //设置最大连接数

	return db
}

func Conn() *sql.DB {
	if global.MysqlAgent == nil {
		global.MysqlAgent = Open()
	}
	return global.MysqlAgent
}
