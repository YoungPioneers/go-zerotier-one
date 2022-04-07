// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ControllerGetMember GET /controller/network/{networkID}member/{memberID}
func (client *Client) ControllerGetMember(networkID string, memberID string) (member Member, err error) {
	resp, err := client.R().Get(fmt.Sprintf("%s/controller/network/%s/member/%s", client.cfg.Address, networkID, memberID))
	if nil != err {
		return member, err
	}

	if resp.IsError() {
		return member, errors.New(resp.Status())
	}

	if err = json.Unmarshal(resp.Body(), &member); nil != err {
		return member, err
	}

	return member, nil
}
