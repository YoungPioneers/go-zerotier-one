// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ControllerStatus .
type ControllerStatus struct {
	Controller bool  `json:"controller"`
	APIVersion uint  `json:"apiVersion"`
	Clock      int64 `json:"clock"`
}

// ControllerStatus GET /controller
func (client *Client) ControllerStatus() (controllerStatus ControllerStatus, err error) {
	resp, err := client.R().Get(fmt.Sprintf("%s/controller", client.cfg.Address))
	if nil != err {
		return controllerStatus, err
	}

	if resp.IsError() {
		return controllerStatus, errors.New(resp.Status())
	}

	if err = json.Unmarshal(resp.Body(), &controllerStatus); nil != err {
		return controllerStatus, err
	}

	return controllerStatus, err
}
