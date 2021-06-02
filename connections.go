package gobudins

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type ConnectionsResponse struct {
	Connections []Connection `json:"connections"`
}

func (ctrl *Controller) GetCurrentUserConnections(token string) (response ConnectionsResponse, err error) {
	return ctrl.GetConnections(UserMe, token)
}

func (ctrl *Controller) GetConnections(userID string, token string) (response ConnectionsResponse, err error) {
	route := fmt.Sprintf("/users/%s/connections", userID)
	err = ctrl.request(http.MethodGet, route, nil, nil, token, &response)
	return response, errors.Wrap(err, "failed to request budget insight api")
}

func (ctrl *Controller) SyncCurrentUserConnection(connectionID, token string) (response Connection, err error) {
	return ctrl.SyncConnection(UserMe, connectionID, token)
}

func (ctrl *Controller) SyncConnection(userID string, connectionID string, token string) (response Connection, err error) {
	route := fmt.Sprintf("/users/%s/connections/%s", userID, connectionID)
	err = ctrl.request(http.MethodPut, route, nil, nil, token, &response)
	return response, errors.Wrap(err, "failed to request budget insight api")
}
