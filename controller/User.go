package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Base struct {
	Name string
	Age  int
	Pass string
}

func (svc *Base) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"massage": "ok"})
}
