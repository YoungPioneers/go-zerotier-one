// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

// Moon .
type Moon struct {
	ID                    string     `json:"id"`
	Roots                 []MoonRoot `json:"roots"`
	Signature             string     `json:"signature"`
	Timestamp             int64      `json:"timestamp"`
	UpdatesMustBeSignedBy string     `json:"updatesMustBeSignedBy"`
	Waiting               bool       `json:"waiting"`
}

// MoonRoot .
type MoonRoot struct {
	Identity        string   `json:"identity"`
	StableEndpoints []string `json:"stableEndpoints"`
}
