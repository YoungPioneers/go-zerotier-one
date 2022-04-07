// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

import (
	"errors"
	"fmt"
)

// DeleteNetwork DELETE /network/{networkID}
func (client *Client) DeleteNetwork(networkID string) (err error) {
	resp, err := client.R().Delete(fmt.Sprintf("%s/network/%s", client.cfg.Address, networkID))
	if nil != err {
		return err
	}

	if resp.IsError() {
		return errors.New(resp.Status())
	}

	return nil
}
