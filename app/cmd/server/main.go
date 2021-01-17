package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fuhiz/docker-go-sample/app/pkg/connecter"
	"github.com/fuhiz/docker-go-sample/app/pkg/controller"
)

func main() {
	// gormのDB接続
	connecter.Setup()

	router := gin.Default()

	// apiの疎通確認用
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Response OK")
	})

	// routing
	r := router.Group("/api/v1")
	controller.Setup(r)

	router.Run(":9000")
}
