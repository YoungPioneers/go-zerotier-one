// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

import (
	"encoding/json"
	"errors"
	"fmt"
)

// GetPeer GET /peer/{address}
func (client *Client) GetPeer(address string) (peer Peer, err error) {
	resp, err := client.R().Get(fmt.Sprintf("%s/peer/%s", client.cfg.Address, address))
	if nil != err {
		return peer, err
	}

	if resp.IsError() {
		return peer, errors.New(resp.Status())
	}

	if err = json.Unmarshal(resp.Body(), &peer); nil != err {
		return peer, err
	}

	return peer, nil
}
