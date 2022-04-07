// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ListNetworks GET /network
func (client *Client) ListNetworks() (networks []Network, err error) {
	resp, err := client.R().Get(fmt.Sprintf("%s/network", client.cfg.Address))
	if nil != err {
		return networks, err
	}

	if resp.IsError() {
		return networks, errors.New(resp.Status())
	}

	networks = make([]Network, 0)
	if err = json.Unmarshal(resp.Body(), &networks); nil != err {
		return networks, err
	}

	return networks, nil
}
