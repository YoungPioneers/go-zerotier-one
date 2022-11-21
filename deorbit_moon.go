// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

import (
	"encoding/json"
	"errors"
	"fmt"
)

// DeorbitMoonOutput .
type DeorbitMoonOutput struct {
	Result bool `json:"result"`
}

// DeorbitMoon POST /moon/{world_id}
func (client *Client) DeorbitMoon(worldID string) (output DeorbitMoonOutput, err error) {
	resp, err := client.R().Delete(fmt.Sprintf("%s/moon/%s", client.cfg.Address, worldID))
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
