// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ControllerListMembers GET /controller/network/{networkID}/member
func (client *Client) ControllerListMembers(networkID string) (members []string, err error) {
	resp, err := client.R().Get(fmt.Sprintf("%s/controller/network/%s/member", client.cfg.Address, networkID))
	if nil != err {
		return members, err
	}

	if resp.IsError() {
		return members, errors.New(resp.Status())
	}

	members = make([]string, 0)
	if err = json.Unmarshal(resp.Body(), &members); nil != err {
		return members, err
	}

	return members, err
}
