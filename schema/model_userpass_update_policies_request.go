// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// UserpassUpdatePoliciesRequest struct for UserpassUpdatePoliciesRequest
type UserpassUpdatePoliciesRequest struct {
	// Use \"token_policies\" instead. If this and \"token_policies\" are both specified, only \"token_policies\" will be used.
	// Deprecated
	Policies []string `json:"policies"`

	// Comma-separated list of policies
	TokenPolicies []string `json:"token_policies"`
}

// NewUserpassUpdatePoliciesRequestWithDefaults instantiates a new UserpassUpdatePoliciesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserpassUpdatePoliciesRequestWithDefaults() *UserpassUpdatePoliciesRequest {
	var this UserpassUpdatePoliciesRequest

	return &this
}

func (o UserpassUpdatePoliciesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["policies"] = o.Policies
	toSerialize["token_policies"] = o.TokenPolicies

	return json.Marshal(toSerialize)
}