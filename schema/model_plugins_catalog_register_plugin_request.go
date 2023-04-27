// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PluginsCatalogRegisterPluginRequest struct for PluginsCatalogRegisterPluginRequest
type PluginsCatalogRegisterPluginRequest struct {
	// The args passed to plugin command.
	Args []string `json:"args"`

	// The command used to start the plugin. The executable defined in this command must exist in vault's plugin directory.
	Command string `json:"command"`

	// The environment variables passed to plugin command. Each entry is of the form \"key=value\".
	Env []string `json:"env"`

	// The SHA256 sum of the executable used in the command field. This should be HEX encoded.
	Sha256 string `json:"sha256"`

	// The type of the plugin, may be auth, secret, or database
	Type string `json:"type"`

	// The semantic version of the plugin to use.
	Version string `json:"version"`
}

// NewPluginsCatalogRegisterPluginRequestWithDefaults instantiates a new PluginsCatalogRegisterPluginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginsCatalogRegisterPluginRequestWithDefaults() *PluginsCatalogRegisterPluginRequest {
	var this PluginsCatalogRegisterPluginRequest

	return &this
}

func (o PluginsCatalogRegisterPluginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["args"] = o.Args
	toSerialize["command"] = o.Command
	toSerialize["env"] = o.Env
	toSerialize["sha256"] = o.Sha256
	toSerialize["type"] = o.Type
	toSerialize["version"] = o.Version

	return json.Marshal(toSerialize)
}