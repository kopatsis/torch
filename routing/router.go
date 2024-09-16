package routing

import (
	"torch/data"

	"github.com/gin-gonic/gin"
)

func New(fullService *data.MainService) *gin.Engine {
	router := gin.Default()
	return router
}

func TestNew() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"ping": "pong"})
	})
	return router
}
