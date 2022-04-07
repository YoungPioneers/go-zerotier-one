// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

// Network .
type Network struct {
	AllowDNS          bool     `json:"allowDNS"`
	AllowDefault      bool     `json:"allowDefault"`
	AllowGlobal       bool     `json:"allowGlobal"`
	AssignedAddresses []string `json:"assignedAddresses"`
	Bridge            bool     `json:"bridge"`
	BroadcastEnabled  bool     `json:"broadcastEnabled"`
	DNS               struct {
		Domain  string   `json:"domain"`
		Servers []string `json:"servers"`
	} `json:"dns"`
	ID                     string `json:"id"`
	MAC                    string `json:"mac"`
	MTU                    uint   `json:"mtu"`
	MulticastSubscriptions []struct {
		ADI uint   `json:"adi"`
		MAC string `json:"mac"`
	} `json:"multicastSubscriptions"`
	Name            string `json:"name"`
	NetconfRevision uint   `json:"netconfRevision"`
	PortDeviceName  string `json:"portDeviceName"`
	PortError       uint   `json:"portError"`
	Routes          []struct {
		Flags  uint   `json:"flags"`
		Metric uint   `json:"metric"`
		Target string `json:"target"`
		Via    string `json:"via"`
	} `json:"routes"`
	Status string `json:"status"`
	Type   string `json:"type"`
}
