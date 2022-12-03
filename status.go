// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Status .
type Status struct {
	Address string `json:"address,omitempty"`
	Clock   int64  `json:"clock,omitempty"`
	Config  struct {
		Settings struct {
			AllowTcpFallbackRelay bool `json:"allowTcpFallbackRelay,omitempty"`
			PortMappingEnabled    bool `json:"portMappingEnabled,omitempty"`
			PrimaryPort           uint `json:"primaryPort,omitempty"`
		} `json:"settings,omitempty"`
	} `json:"config,omitempty"`
	Online               bool   `json:"online,omitempty"`
	PlanetWorldId        int64  `json:"planetWorldId,omitempty"`
	PlanetWorldTimestamp int64  `json:"planetWorldTimestamp,omitempty"`
	PublicIdentity       string `json:"publicIdentity,omitempty"`
	TCPFallbackActive    bool   `json:"tcpFallbackActive,omitempty"`
	Version              string `json:"version,omitempty"`
	VersionBuild         uint   `json:"versionBuild,omitempty"`
	VersionMajor         uint   `json:"versionMajor,omitempty"`
	VersionMinor         uint   `json:"versionMinor,omitempty"`
	VersionRev           uint   `json:"versionRev,omitempty"`
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
