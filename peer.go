// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

// Peer .
type Peer struct {
	Address  string `json:"address"`
	IsBonded bool   `json:"isBonded"`
	Latency  int64  `json:"latency"`
	Paths    []struct {
		Active        bool   `json:"active"`
		Address       string `json:"address"`
		Expired       bool   `json:"expired"`
		LastReceive   int64  `json:"lastReceive"`
		LastSend      int64  `json:"lastSend"`
		Preferred     bool   `json:"preferred"`
		TrustedPathId uint   `json:"trustedPathId"`
	} `json:"paths"`
	Role         string `json:"role"`
	Version      string `json:"version"`
	VersionMajor int    `json:"versionMajor"`
	VersionMinor int    `json:"versionMinor"`
	VersionRev   int    `json:"versionRev"`
}
