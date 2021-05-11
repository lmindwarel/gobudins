package gobudins

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebhooksListeners struct {
	OnUserCreated func(User)
}

func (ctrl *Controller) SetupRoutesGin(e *gin.RouterGroup) {
	e.POST("/user-create", ctrl.whUserCreated)
}

func (ctrl *Controller) whUserCreated(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctrl.listeners.OnUserCreated(user)
}
