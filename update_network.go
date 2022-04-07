// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

import (
	"encoding/json"
	"errors"
	"fmt"
)

// UpdateNetwork POST /network/{networkID} update or join a network
func (client *Client) UpdateNetwork(networkID string, input Network) (output Network, err error) {
	resp, err := client.R().SetBody(input).Post(fmt.Sprintf("%s/network/%s", client.cfg.Address, networkID))
	if nil != err {
		return output, err
	}

	if resp.IsError() {
		return output, errors.New(resp.Status())
	}

	if err = json.Unmarshal(resp.Body(), &output); nil != err {
		return output, err
	}

	return output, nil
}
