// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

import (
	"errors"
	"fmt"
)

// ControllerDeleteMember DELETE /controller/network/{networkID}/member/{memberID} update a member
func (client *Client) ControllerDeleteMember(networkID string, memberID string) (err error) {
	resp, err := client.R().Delete(fmt.Sprintf("%s/controller/network/%s/member/%s", client.cfg.Address, networkID, memberID))
	if nil != err {
		return err
	}

	if resp.IsError() {
		return errors.New(resp.Status())
	}

	return nil
}
