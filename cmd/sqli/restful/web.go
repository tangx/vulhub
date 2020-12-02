package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tangx/vulhub/modules/mysqldb"

	_ "github.com/go-sql-driver/mysql"
)

func GinServer() {
	r := gin.Default()
	r.GET("/v0/:user/:passwd", handlerV0)
	r.GET("/v1/:user/:passwd", handlerV1)
	_ = r.Run(":8080")
}

func handlerV0(c *gin.Context) {
	user := c.Param("user")
	passwd := c.Param("passwd")

	data := getUserByNameV0(user, passwd)
	c.JSON(200, data)
}

func getUserByNameV0(name string, password string) interface{} {

	db := mysqldb.Conn()

	user := new(User)
	sql := fmt.Sprintf(`select name,password from users where name='%s' and password='%s' order by id limit 1;`, name, password)

	// fmt.Println(sql)
	_, err := db.Exec(sql)
	if err != nil {
		logrus.Errorf("get user failed: %s", err.Error())
	}

	return user
}

func handlerV1(c *gin.Context) {
	user := c.Param("user")
	passwd := c.Param("passwd")
	data := getUserByNameV1(user, passwd)
	c.JSON(200, data)
}

func getUserByNameV1(name string, password string) interface{} {

	db := mysqldb.Conn()

	user := new(User)

	sql := `select name,password from users where name=? and password=? order by id limit 1;`
	row := db.QueryRow(sql, name, password)
	if err := row.Scan(&user.Name, &user.Password); err != nil {
		logrus.Errorf("get user failed: %s\n", err.Error())
		return nil
	}

	return user
}
