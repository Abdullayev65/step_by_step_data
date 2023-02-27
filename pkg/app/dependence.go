package app

import (
	hendler "github.com/Abdullayev65/step_by_step_data/pkg/handler"
	"github.com/Abdullayev65/step_by_step_data/pkg/mw"
	"github.com/Abdullayev65/step_by_step_data/pkg/service"
	"github.com/Abdullayev65/step_by_step_data/pkg/utill"
	"time"
)

func (a *App) dependence() {
	db, ctx := a.database()
	token := utill.NewToken("salat", 26*time.Hour)
	s := service.New(db, ctx)
	a.MW = mw.New(token, s)
	a.handler = hendler.New(s, token)

	a.router = a.initRouters()
}
