// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

// ControllerNetwork .
type ControllerNetwork struct {
	ID              string `json:"id"`
	NWID            string `json:"nwid"`
	ObjType         string `json:"objtype"`
	Name            string `json:"name"`
	CreationTime    int64  `json:"creationTime"`
	Private         bool   `json:"private"`
	EnableBroadcast bool   `json:"enableBroadcast"`
	V4AssignMode    struct {
		ZT bool `json:"zt"`
	} `json:"v4AssignMode"`
	V6AssignMode struct {
		SixPlane bool `json:"6plane"`
		RFC4193  bool `json:"rfc4193"`
		ZT       bool `json:"zt"`
	} `json:"v6AssignMode"`
	MTU            uint `json:"mtu"`
	MulticastLimit uint `json:"multicastLimit"`
	Revision       uint `json:"revision"`
	Routes         []struct {
		Target string `json:"target"`
		Via    string `json:"via"`
	} `json:"routes"`
	IPAssignmentPools []struct {
		IPRangeStart string `json:"ipRangeStart"`
		IPRangeEnd   string `json:"ipRangeEnd"`
	}
	Rules []struct {
		Not  bool   `json:"not"`
		Or   bool   `json:"or"`
		Type string `json:"type"`
	} `json:"rules"`
	Capabilities      []interface{} `json:"capabilities"`
	Tags              []interface{} `json:"tags"`
	RemoteTraceTarget string        `json:"remoteTraceTarget"`
	RemoteTraceLevel  uint          `json:"remoteTraceLevel"`
}
