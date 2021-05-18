package gobudins

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func (ctrl *Controller) GetCurrentUserAccountInvestments(accountID string, token string) (accounts []Account, err error) {
	return ctrl.GetAccountInvestments(UserMe, accountID, token)
}

func (ctrl *Controller) GetAccountInvestments(userID string, accountID string, token string) (accounts []Account, err error) {
	route := fmt.Sprintf("/users/%s/accounts/%s/investments", userID, accountID)
	err = ctrl.request(http.MethodGet, route, nil, nil, "", &accounts)

	return accounts, errors.Wrap(err, "failed to request budget insight api")
}
