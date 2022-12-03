// Copyright (c) 2021, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package zerotier

// Moon .
type Moon struct {
	ID                    string     `json:"id,omitempty"`
	ObjType               string     `json:"objtype,omitempty"`
	WorldType             string     `json:"worldType,omitempty"`
	Roots                 []MoonRoot `json:"roots,omitempty"`
	Signature             string     `json:"signature,omitempty"`
	Timestamp             int64      `json:"timestamp,omitempty"`
	Waiting               bool       `json:"waiting,omitempty"`
	SigningKey            string     `json:"signingKey,omitempty"`
	SigningKey_SECRET     string     `json:"signingKey_SECRET,omitempty"`
	UpdatesMustBeSignedBy string     `json:"updatesMustBeSignedBy,omitempty"`
}

// MoonRoot .
type MoonRoot struct {
	Identity        string   `json:"identity,omitempty"`
	StableEndpoints []string `json:"stableEndpoints,omitempty"`
}
