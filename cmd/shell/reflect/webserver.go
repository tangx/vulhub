package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WebServer() {
	r := gin.Default()

	r.POST("/v0/cmd", handler)
	_ = r.Run(":8080")
}

func handler(c *gin.Context) {
	var cmd Cmd

	if err := c.BindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := cmd.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
