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
