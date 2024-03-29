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

// ErrorInner struct for ErrorInner
type ErrorInner struct {
	// The error code relevant to this error
	Code *string `json:"code,omitempty"`
	// A short desription for the error type
	Title *string `json:"title,omitempty"`
	// A detailed message of the error encountered
	Detail *string `json:"detail,omitempty"`
	// The path to the value that caused the error
	Path *string `json:"path,omitempty"`
}

// NewErrorInner instantiates a new ErrorInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorInner() *ErrorInner {
	this := ErrorInner{}
	return &this
}

// NewErrorInnerWithDefaults instantiates a new ErrorInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorInnerWithDefaults() *ErrorInner {
	this := ErrorInner{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ErrorInner) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorInner) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ErrorInner) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ErrorInner) SetCode(v string) {
	o.Code = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ErrorInner) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorInner) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ErrorInner) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ErrorInner) SetTitle(v string) {
	o.Title = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ErrorInner) GetDetail() string {
	if o == nil || isNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorInner) GetDetailOk() (*string, bool) {
	if o == nil || isNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ErrorInner) HasDetail() bool {
	if o != nil && !isNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ErrorInner) SetDetail(v string) {
	o.Detail = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ErrorInner) GetPath() string {
	if o == nil || isNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorInner) GetPathOk() (*string, bool) {
	if o == nil || isNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ErrorInner) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ErrorInner) SetPath(v string) {
	o.Path = &v
}

func (o ErrorInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !isNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableErrorInner struct {
	value *ErrorInner
	isSet bool
}

func (v NullableErrorInner) Get() *ErrorInner {
	return v.value
}

func (v *NullableErrorInner) Set(val *ErrorInner) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorInner) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorInner(val *ErrorInner) *NullableErrorInner {
	return &NullableErrorInner{value: val, isSet: true}
}

func (v NullableErrorInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
