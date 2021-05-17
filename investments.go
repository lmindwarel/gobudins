package gobudins

// func (ctrl *Controller) GetToken(ccd ConnectCallbackData) (token string, err error) {
// 	var authToken GetAccessTokenResponse
// 	err = ctrl.request(http.MethodPost, RouteAccessToken, nil, AskForToken{
// 		Code:         ccd.Code,
// 		ClientID:     ctrl.config.ClientID,
// 		ClientSecret: ctrl.config.ClientSecret,
// 	}, "", &authToken)
// 	if err != nil {
// 		return token, errors.Wrap(err, "failed to request budget insight api")
// 	}

// 	token = authToken.AccessToken

// 	return token, err
// }
