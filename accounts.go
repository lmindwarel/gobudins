package gobudins

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func (ctrl *Controller) GetCurrentUserAccounts(token string) (response AccountsResponse, err error) {
	return ctrl.GetAccounts(UserMe, token)
}

func (ctrl *Controller) GetAccounts(userID string, token string) (response AccountsResponse, err error) {
	route := fmt.Sprintf("%s/%s%s", RouteUsers, userID, RouteAccounts)
	err = ctrl.request(http.MethodGet, route, nil, nil, token, &response)
	return response, errors.Wrap(err, "failed to request budget insight api")
}

func (ctrl *Controller) GetCurrentUserConnectionAccounts(connectionID string, token string) (response AccountsResponse, err error) {
	return ctrl.GetConnectionAccounts(UserMe, connectionID, token)
}

func (ctrl *Controller) GetConnectionAccounts(userID string, connectionID string, token string) (response AccountsResponse, err error) {
	route := fmt.Sprintf("%s/%s/connections/%s%s", RouteUsers, userID, connectionID, RouteAccounts)
	err = ctrl.request(http.MethodGet, route, nil, nil, token, &response)
	return response, errors.Wrap(err, "failed to request budget insight api")
}

func (ctrl *Controller) UpdateCurrentUserAccount(accountID string, update UpdateAccount, token string) error {
	return ctrl.UpdateUserAccount(UserMe, accountID, update, token)
}

func (ctrl *Controller) UpdateUserAccount(userID string, accountID string, update UpdateAccount, token string) error {
	route := fmt.Sprintf("%s/%s%s/%s", RouteUsers, userID, RouteAccounts, accountID)
	return ctrl.request(http.MethodPut, route, nil, update, token, nil)
}
