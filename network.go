// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

// Network .
type Network struct {
	AllowDNS               bool                           `json:"allowDNS,omitempty"`
	AllowDefault           bool                           `json:"allowDefault,omitempty"`
	AllowGlobal            bool                           `json:"allowGlobal,omitempty"`
	AssignedAddresses      []string                       `json:"assignedAddresses,omitempty"`
	Bridge                 bool                           `json:"bridge,omitempty"`
	BroadcastEnabled       bool                           `json:"broadcastEnabled,omitempty"`
	DNS                    NetworkDNS                     `json:"dns,omitempty"`
	ID                     string                         `json:"id,omitempty"`
	MAC                    string                         `json:"mac,omitempty"`
	MTU                    uint                           `json:"mtu,omitempty"`
	MulticastSubscriptions []NetworkMulticastSubscription `json:"multicastSubscriptions,omitempty"`
	Name                   string                         `json:"name,omitempty"`
	NetconfRevision        uint                           `json:"netconfRevision,omitempty"`
	PortDeviceName         string                         `json:"portDeviceName,omitempty"`
	PortError              uint                           `json:"portError,omitempty"`
	Routes                 []NetworkRoute                 `json:"routes,omitempty"`
	Rules                  []NetworkRule                  `json:"rules,omitempty"`
	Status                 string                         `json:"status,omitempty"`
	Type                   string                         `json:"type,omitempty"`
}

// NetworkDNS .
type NetworkDNS struct {
	Domain  string   `json:"domain,omitempty"`
	Servers []string `json:"servers,omitempty"`
}

// NetworkMulticastSubscription .
type NetworkMulticastSubscription struct {
	ADI uint   `json:"adi,omitempty"`
	MAC string `json:"mac,omitempty"`
}

// NetworkRoute .
type NetworkRoute struct {
	Flags  uint   `json:"flags,omitempty"`
	Metric uint   `json:"metric,omitempty"`
	Target string `json:"target,omitempty"`
	Via    string `json:"via,omitempty"`
}

// NetworkRule .
type NetworkRule struct {
	EtherType uint   `json:"etherType,omitempty"`
	Not       bool   `json:"not,omitempty"`
	Or        bool   `json:"or,omitempty"`
	Type      string `json:"type,omitempty"`
}
