/*
Hava

Hava API

API version: 1.1.2
Contact: support@hava.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package havaclient

import (
	"encoding/json"
)

// EnvironmentSummarySourcesInner struct for EnvironmentSummarySourcesInner
type EnvironmentSummarySourcesInner struct {
	// The unique ID of the source
	Id *string `json:"id,omitempty"`
	// The user configured display name of the source
	DisplayName *string `json:"display_name,omitempty"`
}

// NewEnvironmentSummarySourcesInner instantiates a new EnvironmentSummarySourcesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentSummarySourcesInner() *EnvironmentSummarySourcesInner {
	this := EnvironmentSummarySourcesInner{}
	return &this
}

// NewEnvironmentSummarySourcesInnerWithDefaults instantiates a new EnvironmentSummarySourcesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentSummarySourcesInnerWithDefaults() *EnvironmentSummarySourcesInner {
	this := EnvironmentSummarySourcesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvironmentSummarySourcesInner) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSummarySourcesInner) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentSummarySourcesInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EnvironmentSummarySourcesInner) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *EnvironmentSummarySourcesInner) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSummarySourcesInner) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *EnvironmentSummarySourcesInner) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *EnvironmentSummarySourcesInner) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o EnvironmentSummarySourcesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentSummarySourcesInner struct {
	value *EnvironmentSummarySourcesInner
	isSet bool
}

func (v NullableEnvironmentSummarySourcesInner) Get() *EnvironmentSummarySourcesInner {
	return v.value
}

func (v *NullableEnvironmentSummarySourcesInner) Set(val *EnvironmentSummarySourcesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentSummarySourcesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentSummarySourcesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentSummarySourcesInner(val *EnvironmentSummarySourcesInner) *NullableEnvironmentSummarySourcesInner {
	return &NullableEnvironmentSummarySourcesInner{value: val, isSet: true}
}

func (v NullableEnvironmentSummarySourcesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentSummarySourcesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


