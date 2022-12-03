// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

// Member .
type Member struct {
	ID            string   `json:"id,omitempty"`
	Address       string   `json:"address,omitempty"`
	NWID          string   `json:"nwid,omitempty"`
	Authorized    bool     `json:"authorized,omitempty"`
	ActiveBridge  bool     `json:"activeBridge,omitempty"`
	Identity      string   `json:"identity,omitempty"`
	IPAssignments []string `json:"ipAssignments,omitempty"`
	Revision      uint     `json:"revision,omitempty"`
	VMajor        uint     `json:"vMajor,omitempty"`
	VMinor        uint     `json:"vMinor,omitempty"`
	VRev          uint     `json:"vRev,omitempty"`
	VProto        uint     `json:"vProto,omitempty"`
}
