package gobudins

import (
	"net/http"

	"github.com/pkg/errors"
)

func (ctrl *Controller) GetToken(ccd ConnectCallbackData) (token Token, err error) {
	err = ctrl.request(http.MethodPost, RouteAccessToken, nil, AskForToken{
		Code:         ccd.Code,
		ClientID:     ctrl.config.ClientID,
		ClientSecret: ctrl.config.ClientSecret,
	}, "", &token)

	return token, errors.Wrap(err, "failed to request budget insight api")
}
