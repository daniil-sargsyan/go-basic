package app

import (
	"github.com/daniil-sargsyan/go-basic/pkg/config"
	"github.com/daniil-sargsyan/go-basic/pkg/handler"
	"github.com/daniil-sargsyan/go-basic/pkg/router"
	"github.com/daniil-sargsyan/go-basic/pkg/service"
	"github.com/daniil-sargsyan/go-basic/pkg/storage"
	"github.com/gin-gonic/gin"
)

type AppRunner interface {
	Run() error
	GracefulStop() error
}

type app struct {
	cnf *config.Config
	eng *gin.Engine
}

func New(cnf *config.Config) AppRunner {
	repo := storage.NewStorage(cnf.DB)
	srv := service.NewService(repo)
	hndl := handler.NewHandler(srv)

	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	base := engine.Group(cnf.BasePath)
	router.RegisterRoutes(base, hndl)
	return &app{
		eng: engine,
		cnf: cnf,
	}
}

func (a *app) Run() error {
	return a.eng.Run(a.cnf.Address)
}

func (a *app) GracefulStop() error {
	return nil
}
