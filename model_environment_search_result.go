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

// EnvironmentSearchResult struct for EnvironmentSearchResult
type EnvironmentSearchResult struct {
	// The unique ID of the environment
	Id *string `json:"id,omitempty"`
	// The name of the environment
	Name *string `json:"name,omitempty"`
	// The thumbnail of the environment
	Image *string `json:"image,omitempty"`
	// A link to the environment with any matching resources highlighted
	Url *string `json:"url,omitempty"`
}

// NewEnvironmentSearchResult instantiates a new EnvironmentSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentSearchResult() *EnvironmentSearchResult {
	this := EnvironmentSearchResult{}
	return &this
}

// NewEnvironmentSearchResultWithDefaults instantiates a new EnvironmentSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentSearchResultWithDefaults() *EnvironmentSearchResult {
	this := EnvironmentSearchResult{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvironmentSearchResult) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSearchResult) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentSearchResult) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EnvironmentSearchResult) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvironmentSearchResult) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSearchResult) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentSearchResult) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvironmentSearchResult) SetName(v string) {
	o.Name = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *EnvironmentSearchResult) GetImage() string {
	if o == nil || isNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSearchResult) GetImageOk() (*string, bool) {
	if o == nil || isNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *EnvironmentSearchResult) HasImage() bool {
	if o != nil && !isNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *EnvironmentSearchResult) SetImage(v string) {
	o.Image = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *EnvironmentSearchResult) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSearchResult) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *EnvironmentSearchResult) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *EnvironmentSearchResult) SetUrl(v string) {
	o.Url = &v
}

func (o EnvironmentSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentSearchResult struct {
	value *EnvironmentSearchResult
	isSet bool
}

func (v NullableEnvironmentSearchResult) Get() *EnvironmentSearchResult {
	return v.value
}

func (v *NullableEnvironmentSearchResult) Set(val *EnvironmentSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentSearchResult(val *EnvironmentSearchResult) *NullableEnvironmentSearchResult {
	return &NullableEnvironmentSearchResult{value: val, isSet: true}
}

func (v NullableEnvironmentSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
