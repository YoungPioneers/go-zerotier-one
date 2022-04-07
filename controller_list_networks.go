// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ControllerListNetworks GET /controller/network
func (client *Client) ControllerListNetworks() (networks []string, err error) {
	resp, err := client.R().Get(fmt.Sprintf("%s/controller/network", client.cfg.Address))
	if nil != err {
		return networks, err
	}

	if resp.IsError() {
		return networks, errors.New(resp.Status())
	}

	networks = make([]string, 0)
	if err = json.Unmarshal(resp.Body(), &networks); nil != err {
		return networks, err
	}

	return networks, err
}
