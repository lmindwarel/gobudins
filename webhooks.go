package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebhooksListeners struct {
	OnUserCreated func(WebhookUser)
}

func (ctrl *Controller) SetupRoutesGin(e *gin.RouterGroup) {
	e.POST("/user-create", ctrl.UserCreated)
}

func (ctrl *Controller) UserCreated(c *gin.Context) {
	var user WebhookUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctrl.listeners.OnUserCreated(user)
}
