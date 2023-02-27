package hendler

import (
	"github.com/Abdullayev65/step_by_step_data/pkg/ioPut"
	"github.com/Abdullayev65/step_by_step_data/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) ArticlesByStep(c *gin.Context) {
	step := c.GetInt("step")
	articles := h.Service.ArticlesByStep(step)
	c.JSON(200, articles)
}
func (h *Handler) ArticleAccept(c *gin.Context) {
	articleID := c.GetInt("articleID")
	step := c.GetInt("step")
	var commentInput ioPut.CommentInput
	c.Bind(&commentInput)
	if len(commentInput.Comment) < 1 {
		c.String(400, "comment juda qisqa")
		return
	}
	adminID := h.getUserID(c)
	comment := &models.Comment{Comment: commentInput.Comment, ArticleID: articleID,
		Accepted: true, AdminID: adminID, Step: int8(step)}
	err := h.Service.ArticleAccept(articleID, step, comment)
	if err != nil {
		c.String(http.StatusConflict, err.Error())
		return
	}
	c.JSON(200, comment)
}
func (h *Handler) ArticleReject(c *gin.Context) {
	articleID := c.GetInt("articleID")
	step := c.GetInt("step")
	var commentInput ioPut.CommentInput
	c.Bind(&commentInput)
	if len(commentInput.Comment) < 1 {
		c.String(400, "comment juda qisqa")
		return
	}
	adminID := h.getUserID(c)
	comment := &models.Comment{Comment: commentInput.Comment, ArticleID: articleID,
		Accepted: false, AdminID: adminID, Step: int8(step)}
	err := h.Service.ArticleReject(articleID, step, comment)
	if err != nil {
		c.String(http.StatusConflict, err.Error())
		return
	}
	c.JSON(200, comment)
}

func (h *Handler) AllArticles(c *gin.Context) {
	articles := h.Service.AllArticles()
	c.JSON(200, articles)
}
