// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ListMoons GET /moon
func (client *Client) ListMoons() (moons []Moon, err error) {
	resp, err := client.R().Get(fmt.Sprintf("%s/moon", client.cfg.Address))
	if nil != err {
		return moons, err
	}

	if resp.IsError() {
		return moons, errors.New(resp.Status())
	}

	moons = make([]Moon, 0)
	if err = json.Unmarshal(resp.Body(), &moons); nil != err {
		return moons, err
	}

	return moons, nil
}
