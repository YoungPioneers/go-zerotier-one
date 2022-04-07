// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ControllerCreateNetwork POST /controller/network/{controllerID}______ create a network and generate a random network id
func (client *Client) ControllerCreateNetwork(controllerID string, input ControllerNetwork) (output ControllerNetwork, err error) {
	resp, err := client.R().SetBody(input).Post(fmt.Sprintf("%s/controller/network/%s______", client.cfg.Address, controllerID))
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

// ControllerUpsertNetwork POST /controller/network/{networkID} create or update  a network
func (client *Client) ControllerUpsertNetwork(networkID string, input ControllerNetwork) (output ControllerNetwork, err error) {
	resp, err := client.R().SetBody(input).Post(fmt.Sprintf("%s/controller/network/%s", client.cfg.Address, networkID))
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
