package webhooks

import "github.com/gin-gonic/gin"

type Listener struct {
	OnUserCreated func(models.WebhookUser)
}

type Handler struct {
}

func (a *Handler) setupRoutes(e *gin.Engine) {
	e.POST("/user-create", a.UserCreated)
}

func (a *Handler) UserCreated(c *gin.Context) {

}
