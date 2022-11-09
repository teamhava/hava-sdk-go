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

// EnvironmentRevision struct for EnvironmentRevision
type EnvironmentRevision struct {
	// The unique ID of the environment version
	Id *string `json:"id,omitempty"`
	// The start of the date range this version covers
	From *string `json:"from,omitempty"`
	// The end date of the range this version covers. Will be empty if it's the latest version.
	To *string `json:"to,omitempty"`
}

// NewEnvironmentRevision instantiates a new EnvironmentRevision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentRevision() *EnvironmentRevision {
	this := EnvironmentRevision{}
	return &this
}

// NewEnvironmentRevisionWithDefaults instantiates a new EnvironmentRevision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentRevisionWithDefaults() *EnvironmentRevision {
	this := EnvironmentRevision{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvironmentRevision) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentRevision) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentRevision) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EnvironmentRevision) SetId(v string) {
	o.Id = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *EnvironmentRevision) GetFrom() string {
	if o == nil || isNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentRevision) GetFromOk() (*string, bool) {
	if o == nil || isNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *EnvironmentRevision) HasFrom() bool {
	if o != nil && !isNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *EnvironmentRevision) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *EnvironmentRevision) GetTo() string {
	if o == nil || isNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentRevision) GetToOk() (*string, bool) {
	if o == nil || isNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *EnvironmentRevision) HasTo() bool {
	if o != nil && !isNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *EnvironmentRevision) SetTo(v string) {
	o.To = &v
}

func (o EnvironmentRevision) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !isNil(o.To) {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentRevision struct {
	value *EnvironmentRevision
	isSet bool
}

func (v NullableEnvironmentRevision) Get() *EnvironmentRevision {
	return v.value
}

func (v *NullableEnvironmentRevision) Set(val *EnvironmentRevision) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentRevision) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentRevision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentRevision(val *EnvironmentRevision) *NullableEnvironmentRevision {
	return &NullableEnvironmentRevision{value: val, isSet: true}
}

func (v NullableEnvironmentRevision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentRevision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

