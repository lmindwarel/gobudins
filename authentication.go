package gobudins

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

func (ctrl *Controller) GetToken(ccd ConnectCallbackData) (token Token, err error) {
	err = ctrl.request(http.MethodPost, RouteAccessToken, nil, AskForToken{
		Code: ccd.Code,
		APICredentials: APICredentials{
			ClientID:     ctrl.config.ClientID,
			ClientSecret: ctrl.config.ClientSecret,
		},
	}, "", &token)

	return token, errors.Wrap(err, "failed to request budget insight api")
}

func (ctrl *Controller) RenewToken(userID string) (token Token, err error) {
	formattedUserID, err := strconv.Atoi(userID)
	if err != nil {
		return token, errors.Wrap(err, "failed to get user id has string")
	}
	err = ctrl.request(http.MethodPost, "/auth/renew", nil, AskForTokenRenew{
		APICredentials: APICredentials{
			ClientID:     ctrl.config.ClientID,
			ClientSecret: ctrl.config.ClientSecret,
		},
		UserID:         formattedUserID,
		RevokePrevious: true,
	}, "", &token)

	return token, errors.Wrap(err, "failed to request budget insight api")
}
