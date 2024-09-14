package routing

import (
	"torch/data"

	"github.com/gin-gonic/gin"
)

func New(fullService *data.MainService) *gin.Engine {
	router := gin.Default()
	return router
}
