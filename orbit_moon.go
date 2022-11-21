// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

import (
	"encoding/json"
	"errors"
	"fmt"
)

// OrbitMoonInput .
type OrbitMoonInput struct {
	Seed string `json:"seed"`
}

// OrbitMoon POST /moon/{world_id}
func (client *Client) OrbitMoon(worldID string, seed string) (moon Moon, err error) {
	input := OrbitMoonInput{
		Seed: seed,
	}

	resp, err := client.R().SetBody(input).Post(fmt.Sprintf("%s/moon/%s", client.cfg.Address, worldID))
	if nil != err {
		return moon, err
	}

	if resp.IsError() {
		return moon, errors.New(resp.Status())
	}

	if err = json.Unmarshal(resp.Body(), &moon); nil != err {
		return moon, err
	}

	return moon, nil
}
