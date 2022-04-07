// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ControllerUpdateMember POST /controller/network/{networkID}/member/{memberID} update a member
func (client *Client) ControllerUpdateMember(networkID string, memberID string, input Member) (output Member, err error) {
	resp, err := client.R().SetBody(input).Post(fmt.Sprintf("%s/controller/network/%s/member/%s", client.cfg.Address, networkID, memberID))
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
