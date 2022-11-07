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

// EnvironmentsUpdateRequest struct for EnvironmentsUpdateRequest
type EnvironmentsUpdateRequest struct {
	// The updated name for the environment
	Name *string `json:"name,omitempty"`
	// The updated query string for the environment
	Query string `json:"query"`
	// If true the environment will be set to active and display on your environment list
	Complete *bool `json:"complete,omitempty"`
}

// NewEnvironmentsUpdateRequest instantiates a new EnvironmentsUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentsUpdateRequest(query string) *EnvironmentsUpdateRequest {
	this := EnvironmentsUpdateRequest{}
	this.Query = query
	return &this
}

// NewEnvironmentsUpdateRequestWithDefaults instantiates a new EnvironmentsUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentsUpdateRequestWithDefaults() *EnvironmentsUpdateRequest {
	this := EnvironmentsUpdateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvironmentsUpdateRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentsUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentsUpdateRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvironmentsUpdateRequest) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value
func (o *EnvironmentsUpdateRequest) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *EnvironmentsUpdateRequest) GetQueryOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *EnvironmentsUpdateRequest) SetQuery(v string) {
	o.Query = v
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *EnvironmentsUpdateRequest) GetComplete() bool {
	if o == nil || isNil(o.Complete) {
		var ret bool
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentsUpdateRequest) GetCompleteOk() (*bool, bool) {
	if o == nil || isNil(o.Complete) {
		return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *EnvironmentsUpdateRequest) HasComplete() bool {
	if o != nil && !isNil(o.Complete) {
		return true
	}

	return false
}

// SetComplete gets a reference to the given bool and assigns it to the Complete field.
func (o *EnvironmentsUpdateRequest) SetComplete(v bool) {
	o.Complete = &v
}

func (o EnvironmentsUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["query"] = o.Query
	}
	if !isNil(o.Complete) {
		toSerialize["complete"] = o.Complete
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentsUpdateRequest struct {
	value *EnvironmentsUpdateRequest
	isSet bool
}

func (v NullableEnvironmentsUpdateRequest) Get() *EnvironmentsUpdateRequest {
	return v.value
}

func (v *NullableEnvironmentsUpdateRequest) Set(val *EnvironmentsUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentsUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentsUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentsUpdateRequest(val *EnvironmentsUpdateRequest) *NullableEnvironmentsUpdateRequest {
	return &NullableEnvironmentsUpdateRequest{value: val, isSet: true}
}

func (v NullableEnvironmentsUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentsUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


