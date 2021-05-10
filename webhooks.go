package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lmindwarel/budget-insight-go/models"
)

type WebhooksListeners struct {
	OnUserCreated func(WebhookUser)
}

func (ctrl *Controller) SetupRoutesGin(e *gin.RouterGroup) {
	e.POST("/user-create", ctrl.UserCreated)
}

func (ctrl *Controller) UserCreated(c *gin.Context) {
	var user models.WebhookUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctrl.WebhooksListeners.OnUserCreated(user)
}
