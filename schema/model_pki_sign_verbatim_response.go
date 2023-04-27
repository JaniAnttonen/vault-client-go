// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiSignVerbatimResponse struct for PkiSignVerbatimResponse
type PkiSignVerbatimResponse struct {
	// Certificate Chain
	CaChain []string `json:"ca_chain"`

	// Certificate
	Certificate string `json:"certificate"`

	// Time of expiration
	Expiration string `json:"expiration"`

	// Issuing Certificate Authority
	IssuingCa string `json:"issuing_ca"`

	// Private key
	PrivateKey string `json:"private_key"`

	// Private key type
	PrivateKeyType string `json:"private_key_type"`

	// Serial Number
	SerialNumber string `json:"serial_number"`
}

// NewPkiSignVerbatimResponseWithDefaults instantiates a new PkiSignVerbatimResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiSignVerbatimResponseWithDefaults() *PkiSignVerbatimResponse {
	var this PkiSignVerbatimResponse

	return &this
}

func (o PkiSignVerbatimResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["ca_chain"] = o.CaChain
	toSerialize["certificate"] = o.Certificate
	toSerialize["expiration"] = o.Expiration
	toSerialize["issuing_ca"] = o.IssuingCa
	toSerialize["private_key"] = o.PrivateKey
	toSerialize["private_key_type"] = o.PrivateKeyType
	toSerialize["serial_number"] = o.SerialNumber

	return json.Marshal(toSerialize)
}