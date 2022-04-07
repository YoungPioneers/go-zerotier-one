// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

import (
	"encoding/json"
	"errors"
	"fmt"
)

// GetNetwork GET /network/{networkID}
func (client *Client) GetNetwork(networkID string) (network Network, err error) {
	resp, err := client.R().Get(fmt.Sprintf("%s/network/%s", client.cfg.Address, networkID))
	if nil != err {
		return network, err
	}

	if resp.IsError() {
		return network, errors.New(resp.Status())
	}

	if err = json.Unmarshal(resp.Body(), &network); nil != err {
		return network, err
	}

	return network, nil
}
