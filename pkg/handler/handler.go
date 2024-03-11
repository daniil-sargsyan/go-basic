package handler

import (
	// import "github.com/gin-gonic/gin"

	"github.com/daniil-sargsyan/go-basic/pkg/service"
)

type Handler interface {
	//add methods to handler and implement
	//ex. func myFunc(c *gin.Context)
}

type handler struct {
	srv service.Service
}

func NewHandler(srv service.Service) Handler {
	return &handler{srv: srv}
}
