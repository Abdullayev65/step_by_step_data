package hendler

import (
	"github.com/Abdullayev65/step_by_step_data/pkg/ioPut"
	"github.com/Abdullayev65/step_by_step_data/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ArticleAdd(c *gin.Context) {
	userID := h.getUserID(c)
	var articleInput ioPut.ArticleInput
	c.Bind(&articleInput)
	article := &models.Article{UserID: userID, Data: articleInput.Data, Step: 1, Active: true}
	err := h.Service.ArticleAdd(article)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, article)
}
func (h *Handler) ArticleByID(c *gin.Context) {
	id := c.GetInt("id")
	article, err := h.Service.ArticleGet(id)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, article)
}
func (h *Handler) ArticleAllByUserID(c *gin.Context) {
	userId := h.getUserID(c)
	articles := h.Service.ArticlesByUserID(userId)
	c.JSON(200, articles)
}
