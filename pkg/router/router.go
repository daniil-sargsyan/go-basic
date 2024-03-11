package router

import (
	"github.com/gin-gonic/gin"

	"github.com/daniil-sargsyan/go-basic/pkg/handler"
)

func RegisterRoutes(engine *gin.RouterGroup, hndl handler.Handler) {
	// engine.GET("/", <handler func>)
}
