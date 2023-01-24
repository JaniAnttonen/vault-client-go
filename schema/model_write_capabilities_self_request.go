/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteCapabilitiesSelfRequest struct for WriteCapabilitiesSelfRequest
type WriteCapabilitiesSelfRequest struct {
	// Use 'paths' instead.
	// Deprecated
	Path []string `json:"path"`
	// Paths on which capabilities are being queried.
	Paths []string `json:"paths"`
	// Token for which capabilities are being queried.
	Token string `json:"token"`
}

// NewWriteCapabilitiesSelfRequestWithDefaults instantiates a new WriteCapabilitiesSelfRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteCapabilitiesSelfRequestWithDefaults() *WriteCapabilitiesSelfRequest {
	var this WriteCapabilitiesSelfRequest

	return &this
}

func (o WriteCapabilitiesSelfRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["path"] = o.Path
	toSerialize["paths"] = o.Paths
	toSerialize["token"] = o.Token

	return json.Marshal(toSerialize)
}