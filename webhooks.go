package main

import "github.com/gin-gonic/gin"

type Listener struct {
	OnUserCreated func(WebhookUser)
}

func (a *Controller) setupRoutes(e *gin.RouterGroup) {
	e.POST("/user-create", a.UserCreated)
}

func (a *Controller) UserCreated(c *gin.Context) {

}
