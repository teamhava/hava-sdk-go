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

// EnvironmentFacet The facet contains the resources that are projected to the views.
type EnvironmentFacet struct {
	// The unique ID of the facet
	Id *string `json:"id,omitempty"`
	// All the resources that are included in the environment. Resources can return different attributes based on their type.
	Resources []Resource `json:"resources,omitempty"`
}

// NewEnvironmentFacet instantiates a new EnvironmentFacet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentFacet() *EnvironmentFacet {
	this := EnvironmentFacet{}
	return &this
}

// NewEnvironmentFacetWithDefaults instantiates a new EnvironmentFacet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentFacetWithDefaults() *EnvironmentFacet {
	this := EnvironmentFacet{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvironmentFacet) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentFacet) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentFacet) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EnvironmentFacet) SetId(v string) {
	o.Id = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *EnvironmentFacet) GetResources() []Resource {
	if o == nil || isNil(o.Resources) {
		var ret []Resource
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentFacet) GetResourcesOk() ([]Resource, bool) {
	if o == nil || isNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *EnvironmentFacet) HasResources() bool {
	if o != nil && !isNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []Resource and assigns it to the Resources field.
func (o *EnvironmentFacet) SetResources(v []Resource) {
	o.Resources = v
}

func (o EnvironmentFacet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentFacet struct {
	value *EnvironmentFacet
	isSet bool
}

func (v NullableEnvironmentFacet) Get() *EnvironmentFacet {
	return v.value
}

func (v *NullableEnvironmentFacet) Set(val *EnvironmentFacet) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentFacet) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentFacet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentFacet(val *EnvironmentFacet) *NullableEnvironmentFacet {
	return &NullableEnvironmentFacet{value: val, isSet: true}
}

func (v NullableEnvironmentFacet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentFacet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
