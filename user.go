package gobudins

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func (ctrl *Controller) GetUser(id string, token string) (user User, err error) {
	err = ctrl.request(http.MethodGet, fmt.Sprintf("%s/%s", RouteUsers, id), nil, nil, token, &user)
	if err != nil {
		return user, errors.Wrap(err, "failed to request budget insight api")
	}

	return
}

func (ctrl *Controller) GetTemporaryCode(token string) (code string, err error) {
	err = ctrl.request(http.MethodGet, RouteAuthTokenCode, nil, nil, token, &code)
	if err != nil {
		return code, errors.Wrap(err, "failed to request budget insight api")
	}

	return
}
