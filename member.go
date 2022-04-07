// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

// Member .
type Member struct {
	ID            string   `json:"id"`
	Address       string   `json:"address"`
	NWID          string   `json:"nwid"`
	Authorized    bool     `json:"authorized"`
	ActiveBridge  bool     `json:"activeBridge"`
	Identity      string   `json:"identity"`
	IPAssignments []string `json:"ipAssignments"`
	Revision      uint     `json:"revision"`
	VMajor        uint     `json:"vMajor"`
	VMinor        uint     `json:"vMinor"`
	VRev          uint     `json:"vRev"`
	VProto        uint     `json:"vProto"`
}
