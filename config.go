// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

const (
	// DefaultAddress default zerotier-one address
	DefaultAddress = "http://127.0.0.1:9993"
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
