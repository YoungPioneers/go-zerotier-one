// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Status .
type Status struct {
	Address string `json:"address"`
	Clock   int64  `json:"clock"`
	Config  struct {
		Settings struct {
			AllowTcpFallbackRelay bool `json:"allowTcpFallbackRelay"`
			PortMappingEnabled    bool `json:"portMappingEnabled"`
			PrimaryPort           uint `json:"primaryPort"`
		} `json:"settings"`
	} `json:"config"`
	Online               bool   `json:"online"`
	PlanetWorldId        int64  `json:"planetWorldId"`
	PlanetWorldTimestamp int64  `json:"planetWorldTimestamp"`
	PublicIdentity       string `json:"publicIdentity"`
	TCPFallbackActive    bool   `json:"tcpFallbackActive"`
	Version              string `json:"version"`
	VersionBuild         uint   `json:"versionBuild"`
	VersionMajor         uint   `json:"versionMajor"`
	VersionMinor         uint   `json:"versionMinor"`
	VersionRev           uint   `json:"versionRev"`
}

// Status GET /status
func (client *Client) Status() (status Status, err error) {
	resp, err := client.R().Get(fmt.Sprintf("%s/status", client.cfg.Address))
	if nil != err {
		return status, err
	}

	if resp.IsError() {
		return status, errors.New(resp.Status())
	}

	if err = json.Unmarshal(resp.Body(), &status); nil != err {
		return status, err
	}

	return status, err
}
