package app

import (
	hendler "github.com/Abdullayev65/step_by_step_data/pkg/handler"
	"github.com/Abdullayev65/step_by_step_data/pkg/mw"
	"github.com/gin-gonic/gin"
)

type App struct {
	handler *hendler.Handler
	router  *gin.Engine
	MW      *mw.MW
}

func New() *App {
	a := &App{}
	a.dependence()
	return a
}

func (a *App) Run() {
	err := a.router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
