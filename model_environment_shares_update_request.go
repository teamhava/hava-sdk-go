/*
Hava

Hava API

API version: 1.1.3
Contact: support@hava.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package havaclient

import (
	"encoding/json"
)

// EnvironmentSharesUpdateRequest struct for EnvironmentSharesUpdateRequest
type EnvironmentSharesUpdateRequest struct {
	Config *EnvironmentSharesCreateRequestConfig `json:"config,omitempty"`
}

// NewEnvironmentSharesUpdateRequest instantiates a new EnvironmentSharesUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentSharesUpdateRequest() *EnvironmentSharesUpdateRequest {
	this := EnvironmentSharesUpdateRequest{}
	return &this
}

// NewEnvironmentSharesUpdateRequestWithDefaults instantiates a new EnvironmentSharesUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentSharesUpdateRequestWithDefaults() *EnvironmentSharesUpdateRequest {
	this := EnvironmentSharesUpdateRequest{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *EnvironmentSharesUpdateRequest) GetConfig() EnvironmentSharesCreateRequestConfig {
	if o == nil || isNil(o.Config) {
		var ret EnvironmentSharesCreateRequestConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSharesUpdateRequest) GetConfigOk() (*EnvironmentSharesCreateRequestConfig, bool) {
	if o == nil || isNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *EnvironmentSharesUpdateRequest) HasConfig() bool {
	if o != nil && !isNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given EnvironmentSharesCreateRequestConfig and assigns it to the Config field.
func (o *EnvironmentSharesUpdateRequest) SetConfig(v EnvironmentSharesCreateRequestConfig) {
	o.Config = &v
}

func (o EnvironmentSharesUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentSharesUpdateRequest struct {
	value *EnvironmentSharesUpdateRequest
	isSet bool
}

func (v NullableEnvironmentSharesUpdateRequest) Get() *EnvironmentSharesUpdateRequest {
	return v.value
}

func (v *NullableEnvironmentSharesUpdateRequest) Set(val *EnvironmentSharesUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentSharesUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentSharesUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentSharesUpdateRequest(val *EnvironmentSharesUpdateRequest) *NullableEnvironmentSharesUpdateRequest {
	return &NullableEnvironmentSharesUpdateRequest{value: val, isSet: true}
}

func (v NullableEnvironmentSharesUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentSharesUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
