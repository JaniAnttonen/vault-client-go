// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AwsGenerateCredentialsWithParametersRequest struct for AwsGenerateCredentialsWithParametersRequest
type AwsGenerateCredentialsWithParametersRequest struct {
	// ARN of role to assume when credential_type is assumed_role
	RoleArn string `json:"role_arn,omitempty"`

	// Session name to use when assuming role. Max chars: 64
	RoleSessionName string `json:"role_session_name,omitempty"`

	// Lifetime of the returned credentials in seconds
	Ttl string `json:"ttl,omitempty"`
}

// NewAwsGenerateCredentialsWithParametersRequestWithDefaults instantiates a new AwsGenerateCredentialsWithParametersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsGenerateCredentialsWithParametersRequestWithDefaults() *AwsGenerateCredentialsWithParametersRequest {
	var this AwsGenerateCredentialsWithParametersRequest

	this.Ttl = "3600"

	return &this
}