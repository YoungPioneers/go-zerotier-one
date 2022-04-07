// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.
// https://docs.zerotier.com/service/v1

package zerotier

import (
	"os"

	"github.com/go-resty/resty/v2"
)

// Client .
type Client struct {
	cfg Config

	// http client
	*resty.Client
}

// New Initialize new zerotier client from confi
func New(cfg Config) (client *Client, err error) {
	if "" == cfg.Address {
		cfg.Address = DefaultAddress
	}
	// if api key is not provided, try to find it in the default path
	if "" == cfg.Auth {
		bs, err := os.ReadFile("/var/lib/zerotier-one/authtoken.secret")
		if nil != err {
			return nil, err
		}

		cfg.Auth = string(bs)
	}

	client = &Client{
		cfg:    cfg,
		Client: resty.New(),
	}

	client.Client = client.OnBeforeRequest(func(restyClient *resty.Client, restyRequest *resty.Request) error {
		restyRequest.SetHeader("X-ZT1-Auth", cfg.Auth)

		return nil
	})

	return client, nil
}
