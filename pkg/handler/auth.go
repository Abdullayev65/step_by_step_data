package hendler

import (
	"github.com/Abdullayev65/step_by_step_data/pkg/ioPut"
	"github.com/Abdullayev65/step_by_step_data/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	var sign ioPut.Sign
	c.Bind(&sign)
	user := models.User{Username: sign.Username, Password: sign.Password, Admin: false}
	err := h.Service.UserAdd(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, &user)
}
func (h *Handler) SignIn(c *gin.Context) {
	var sign ioPut.Sign
	c.Bind(&sign)
	user, err := h.Service.UserByUsername(sign.Username)
	if err != nil || user.Password != sign.Password {
		c.String(http.StatusBadRequest, "username or password wrong")
		return
	}
	token, err := h.TokenJWT.Generate(user.ID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
