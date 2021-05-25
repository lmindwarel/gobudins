package gobudins

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func (ctrl *Controller) GetCurrentUserAccountInvestments(accountID string, token string) (response InvestmentsResponse, err error) {
	return ctrl.GetAccountInvestments(UserMe, accountID, token)
}

func (ctrl *Controller) GetAccountInvestments(userID string, accountID string, token string) (response InvestmentsResponse, err error) {
	route := fmt.Sprintf("/users/%s/accounts/%s/investments", userID, accountID)
	err = ctrl.request(http.MethodGet, route, nil, nil, token, &response)
	return response, errors.Wrap(err, "failed to request budget insight api")
}

func (ctrl *Controller) GetCurrentUserInvestments(token string) (InvestmentsResponse, error) {
	return ctrl.GetInvestments(UserMe, token)
}

func (ctrl *Controller) GetInvestments(userID string, token string) (response InvestmentsResponse, err error) {
	route := fmt.Sprintf("/users/%s/investments", userID)
	err = ctrl.request(http.MethodGet, route, nil, nil, token, &response)
	return response, errors.Wrap(err, "failed to request budget insight api")
}
