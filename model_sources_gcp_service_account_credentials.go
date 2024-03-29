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

// SourcesGCPServiceAccountCredentials The parameters required to create a GCP Service Account Credentials Source
type SourcesGCPServiceAccountCredentials struct {
	Name *string `json:"name,omitempty"`
	// Must be set to GCP::ServiceAccountCredentials
	Type *string `json:"type,omitempty"`
	// Base64 encoded credentials file
	EncodedFile *string `json:"encoded_file,omitempty"`
}

// NewSourcesGCPServiceAccountCredentials instantiates a new SourcesGCPServiceAccountCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourcesGCPServiceAccountCredentials() *SourcesGCPServiceAccountCredentials {
	this := SourcesGCPServiceAccountCredentials{}
	return &this
}

// NewSourcesGCPServiceAccountCredentialsWithDefaults instantiates a new SourcesGCPServiceAccountCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourcesGCPServiceAccountCredentialsWithDefaults() *SourcesGCPServiceAccountCredentials {
	this := SourcesGCPServiceAccountCredentials{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SourcesGCPServiceAccountCredentials) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesGCPServiceAccountCredentials) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SourcesGCPServiceAccountCredentials) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SourcesGCPServiceAccountCredentials) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SourcesGCPServiceAccountCredentials) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesGCPServiceAccountCredentials) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SourcesGCPServiceAccountCredentials) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SourcesGCPServiceAccountCredentials) SetType(v string) {
	o.Type = &v
}

// GetEncodedFile returns the EncodedFile field value if set, zero value otherwise.
func (o *SourcesGCPServiceAccountCredentials) GetEncodedFile() string {
	if o == nil || isNil(o.EncodedFile) {
		var ret string
		return ret
	}
	return *o.EncodedFile
}

// GetEncodedFileOk returns a tuple with the EncodedFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesGCPServiceAccountCredentials) GetEncodedFileOk() (*string, bool) {
	if o == nil || isNil(o.EncodedFile) {
		return nil, false
	}
	return o.EncodedFile, true
}

// HasEncodedFile returns a boolean if a field has been set.
func (o *SourcesGCPServiceAccountCredentials) HasEncodedFile() bool {
	if o != nil && !isNil(o.EncodedFile) {
		return true
	}

	return false
}

// SetEncodedFile gets a reference to the given string and assigns it to the EncodedFile field.
func (o *SourcesGCPServiceAccountCredentials) SetEncodedFile(v string) {
	o.EncodedFile = &v
}

func (o SourcesGCPServiceAccountCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.EncodedFile) {
		toSerialize["encoded_file"] = o.EncodedFile
	}
	return json.Marshal(toSerialize)
}

type NullableSourcesGCPServiceAccountCredentials struct {
	value *SourcesGCPServiceAccountCredentials
	isSet bool
}

func (v NullableSourcesGCPServiceAccountCredentials) Get() *SourcesGCPServiceAccountCredentials {
	return v.value
}

func (v *NullableSourcesGCPServiceAccountCredentials) Set(val *SourcesGCPServiceAccountCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableSourcesGCPServiceAccountCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableSourcesGCPServiceAccountCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourcesGCPServiceAccountCredentials(val *SourcesGCPServiceAccountCredentials) *NullableSourcesGCPServiceAccountCredentials {
	return &NullableSourcesGCPServiceAccountCredentials{value: val, isSet: true}
}

func (v NullableSourcesGCPServiceAccountCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourcesGCPServiceAccountCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
