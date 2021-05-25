package gobudins

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebhooksListeners struct {
	OnUserCreated   func(User)
	OnAccountSynced func(SyncedAccount)
}

func (ctrl *Controller) SetupRoutesGin(e *gin.RouterGroup) {
	e.POST("/user/created", ctrl.whUserCreated)
	e.POST("/accounts/synced", ctrl.whAccountSynced)
}

func (ctrl *Controller) whUserCreated(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctrl.listeners.OnUserCreated(user)
}

func (ctrl *Controller) whAccountSynced(c *gin.Context) {
	var account SyncedAccount
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctrl.listeners.OnAccountSynced(account)
}
