/*
Hava

Hava API

API version: 1.1.1
Contact: support@hava.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package havaclient

import (
	"encoding/json"
)

// EnvironmentsBody struct for EnvironmentsBody
type EnvironmentsBody struct {
	// The name for your environment
	Name *string `json:"name,omitempty"`
	// The query string used to create your environment
	Query string `json:"query"`
}

// NewEnvironmentsBody instantiates a new EnvironmentsBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentsBody(query string) *EnvironmentsBody {
	this := EnvironmentsBody{}
	this.Query = query
	return &this
}

// NewEnvironmentsBodyWithDefaults instantiates a new EnvironmentsBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentsBodyWithDefaults() *EnvironmentsBody {
	this := EnvironmentsBody{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvironmentsBody) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentsBody) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentsBody) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvironmentsBody) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value
func (o *EnvironmentsBody) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *EnvironmentsBody) GetQueryOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *EnvironmentsBody) SetQuery(v string) {
	o.Query = v
}

func (o EnvironmentsBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentsBody struct {
	value *EnvironmentsBody
	isSet bool
}

func (v NullableEnvironmentsBody) Get() *EnvironmentsBody {
	return v.value
}

func (v *NullableEnvironmentsBody) Set(val *EnvironmentsBody) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentsBody) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentsBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentsBody(val *EnvironmentsBody) *NullableEnvironmentsBody {
	return &NullableEnvironmentsBody{value: val, isSet: true}
}

func (v NullableEnvironmentsBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentsBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


