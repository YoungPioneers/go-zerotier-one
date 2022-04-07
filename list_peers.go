// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ListPeers GET /peer
func (client *Client) ListPeers() (peers []Peer, err error) {
	resp, err := client.R().Get(fmt.Sprintf("%s/peer", client.cfg.Address))
	if nil != err {
		return peers, err
	}

	if resp.IsError() {
		return peers, errors.New(resp.Status())
	}

	peers = make([]Peer, 0)
	if err = json.Unmarshal(resp.Body(), &peers); nil != err {
		return peers, err
	}

	return peers, nil
}
