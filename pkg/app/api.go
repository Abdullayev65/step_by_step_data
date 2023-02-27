package app

import "github.com/gin-gonic/gin"

func (a *App) initRouters() *gin.Engine {
	router := gin.Default()

	user := router.Group("/auth")
	{
		user.POST("/sign-up", a.handler.SignUp)
		user.POST("/sign-in", a.handler.SignIn)
	}
	article := router.Group("/article")
	{
		article.POST("/",
			a.MW.UserIDFromToken, a.handler.ArticleAdd)
		article.GET("/:id",
			a.MW.SetIntFromParam("id"), a.handler.ArticleByID)
		article.GET("/all",
			a.MW.UserIDFromToken, a.handler.ArticleAllByUserID)
	}
	admin := router.Group("/admin", a.MW.UserIDFromToken, a.MW.CheckAdmin)
	{
		admin.GET("/articles/:step",
			a.MW.SetIntFromParam("step"), a.handler.ArticlesByStep)
		admin.GET("/all-articles", a.handler.AllArticles)
		admin.POST("/article/accept",
			a.MW.SetIntFromQuery("articleID", "step"), a.handler.ArticleAccept)
		admin.POST("/article/reject",
			a.MW.SetIntFromQuery("articleID", "step"), a.handler.ArticleReject)
	}
	return router
}
