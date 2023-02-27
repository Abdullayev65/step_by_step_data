package hendler

import (
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
	err := h.Service.ArticleAccept(articleID, step)
	if err != nil {
		c.String(http.StatusConflict, err.Error())
		return
	}
	c.String(200, "DONE")
}
func (h *Handler) ArticleReject(c *gin.Context) {
	articleID := c.GetInt("articleID")
	err := h.Service.ArticleReject(articleID)
	if err != nil {
		c.String(http.StatusConflict, err.Error())
		return
	}
	c.String(200, "DONE")
}

func (h *Handler) AllArticles(c *gin.Context) {
	articles := h.Service.AllArticles()
	c.JSON(200, articles)
}
