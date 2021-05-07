package main

import (
	"net/http"

	"github.com/lmindwarel/models"
	"github.com/pkg/errors"
)

func (ctrl *Controller) GetToken(ccd models.ConnectCallbackData) (token string, err error) {
	var authToken models.GetAccessTokenResponse
	err = ctrl.request(http.MethodPost, models.RouteAccessToken, nil, models.AskForToken{
		Code:         ccd.Code,
		ClientID:     ctrl.config.ClientID,
		ClientSecret: ctrl.config.ClientSecret,
	}, "", &authToken)
	if err != nil {
		return token, errors.Wrap(err, "failed to request budget insight api")
	}

	token = authToken.AccessToken

	return token, err
}
