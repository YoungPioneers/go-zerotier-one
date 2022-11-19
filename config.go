// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

const (
	// DefaultAddress default zerotier-one address
	DefaultAddress = "http://127.0.0.1:9993"
	// LocalConfigPath zerotier local.conf default path
	LocalConfigPath = "/var/lib/zerotier-one/local.conf"
)

// Config .
type Config struct {
	// Auth API Key for zerotier-one
	Auth    string `json:"auth" yaml:"auth"`
	Address string `json:"address" yaml:"address"`
}

// WithAuth .
func (cfg *Config) WithAuth(auth string) *Config {
	cfg.Auth = auth

	return cfg
}

// WithAddress .
func (cfg *Config) WithAddress(address string) *Config {
	cfg.Address = address

	return cfg
}

// LocalConfig .
type LocalConfig struct {
	Physical map[string]NetworkConfig `json:"physical,omitempty"`
	Virtual  map[string]VirtualConfig `json:"virtual,omitempty"`
	Settings SettingsConfig           `json:"settings,omitempty"`
}

// NetworkConfig .
type NetworkConfig struct {
	BlackList     bool  `json:"blacklist,omitempty"`
	TrustedPathID int64 `json:"trustedPathId,omitempty"`
	MTU           int64 `json:"mtu,omitempty"`
}

// VirtualConfig .
type VirtualConfig struct {
	Try       []string `json:"try,omitempty"`
	BlackList []string `json:"blacklist,omitempty"`
}

// SettingsConfig .
type SettingsConfig struct {
	PrimaryPort              int64    `json:"primaryPort,omitempty"`
	SecondaryPort            int64    `json:"secondaryPort,omitempty"`
	TertiaryPort             int64    `json:"tertiaryPort,omitempty"`
	PortMappingEnabled       bool     `json:"portMappingEnabled,omitempty"`
	AllowSecondaryPort       bool     `json:"allowSecondaryPort,omitempty"`
	SoftwareUpdate           string   `json:"softwareUpdate,omitempty"`
	SoftwareUpdateChannel    string   `json:"softwareUpdateChannel,omitempty"`
	SoftwareUpdateDist       bool     `json:"softwareUpdateDist,omitempty"`
	InterfacePrefixBlacklist []string `json:"interfacePrefixBlacklist,omitempty"`
	AllowManagementFrom      []string `json:"allowManagementFrom,omitempty"`
	Bind                     []string `json:"bind,omitempty"`
	AllowTCPFallbackRelay    bool     `json:"allowTcpFallbackRelay,omitempty"`
	MultipathMode            int64    `json:"multipathMode,omitempty"`
}
