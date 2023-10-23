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

// EnvironmentLatestRevisions A list of the latest versions of this environment
type EnvironmentLatestRevisions struct {
	// The amount of results per page
	LimitValue *int32 `json:"limit_value,omitempty"`
	// The current page
	CurrentPage *int32 `json:"current_page,omitempty"`
	// The total pages, based on limit and total versions
	TotalPages *int32 `json:"total_pages,omitempty"`
	// The total versions available
	TotalCount *int32                `json:"total_count,omitempty"`
	Map        []EnvironmentRevision `json:"map,omitempty"`
}

// NewEnvironmentLatestRevisions instantiates a new EnvironmentLatestRevisions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentLatestRevisions() *EnvironmentLatestRevisions {
	this := EnvironmentLatestRevisions{}
	return &this
}

// NewEnvironmentLatestRevisionsWithDefaults instantiates a new EnvironmentLatestRevisions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentLatestRevisionsWithDefaults() *EnvironmentLatestRevisions {
	this := EnvironmentLatestRevisions{}
	return &this
}

// GetLimitValue returns the LimitValue field value if set, zero value otherwise.
func (o *EnvironmentLatestRevisions) GetLimitValue() int32 {
	if o == nil || isNil(o.LimitValue) {
		var ret int32
		return ret
	}
	return *o.LimitValue
}

// GetLimitValueOk returns a tuple with the LimitValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLatestRevisions) GetLimitValueOk() (*int32, bool) {
	if o == nil || isNil(o.LimitValue) {
		return nil, false
	}
	return o.LimitValue, true
}

// HasLimitValue returns a boolean if a field has been set.
func (o *EnvironmentLatestRevisions) HasLimitValue() bool {
	if o != nil && !isNil(o.LimitValue) {
		return true
	}

	return false
}

// SetLimitValue gets a reference to the given int32 and assigns it to the LimitValue field.
func (o *EnvironmentLatestRevisions) SetLimitValue(v int32) {
	o.LimitValue = &v
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise.
func (o *EnvironmentLatestRevisions) GetCurrentPage() int32 {
	if o == nil || isNil(o.CurrentPage) {
		var ret int32
		return ret
	}
	return *o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLatestRevisions) GetCurrentPageOk() (*int32, bool) {
	if o == nil || isNil(o.CurrentPage) {
		return nil, false
	}
	return o.CurrentPage, true
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *EnvironmentLatestRevisions) HasCurrentPage() bool {
	if o != nil && !isNil(o.CurrentPage) {
		return true
	}

	return false
}

// SetCurrentPage gets a reference to the given int32 and assigns it to the CurrentPage field.
func (o *EnvironmentLatestRevisions) SetCurrentPage(v int32) {
	o.CurrentPage = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *EnvironmentLatestRevisions) GetTotalPages() int32 {
	if o == nil || isNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLatestRevisions) GetTotalPagesOk() (*int32, bool) {
	if o == nil || isNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *EnvironmentLatestRevisions) HasTotalPages() bool {
	if o != nil && !isNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *EnvironmentLatestRevisions) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *EnvironmentLatestRevisions) GetTotalCount() int32 {
	if o == nil || isNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLatestRevisions) GetTotalCountOk() (*int32, bool) {
	if o == nil || isNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *EnvironmentLatestRevisions) HasTotalCount() bool {
	if o != nil && !isNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *EnvironmentLatestRevisions) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetMap returns the Map field value if set, zero value otherwise.
func (o *EnvironmentLatestRevisions) GetMap() []EnvironmentRevision {
	if o == nil || isNil(o.Map) {
		var ret []EnvironmentRevision
		return ret
	}
	return o.Map
}

// GetMapOk returns a tuple with the Map field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLatestRevisions) GetMapOk() ([]EnvironmentRevision, bool) {
	if o == nil || isNil(o.Map) {
		return nil, false
	}
	return o.Map, true
}

// HasMap returns a boolean if a field has been set.
func (o *EnvironmentLatestRevisions) HasMap() bool {
	if o != nil && !isNil(o.Map) {
		return true
	}

	return false
}

// SetMap gets a reference to the given []EnvironmentRevision and assigns it to the Map field.
func (o *EnvironmentLatestRevisions) SetMap(v []EnvironmentRevision) {
	o.Map = v
}

func (o EnvironmentLatestRevisions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LimitValue) {
		toSerialize["limit_value"] = o.LimitValue
	}
	if !isNil(o.CurrentPage) {
		toSerialize["current_page"] = o.CurrentPage
	}
	if !isNil(o.TotalPages) {
		toSerialize["total_pages"] = o.TotalPages
	}
	if !isNil(o.TotalCount) {
		toSerialize["total_count"] = o.TotalCount
	}
	if !isNil(o.Map) {
		toSerialize["map"] = o.Map
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentLatestRevisions struct {
	value *EnvironmentLatestRevisions
	isSet bool
}

func (v NullableEnvironmentLatestRevisions) Get() *EnvironmentLatestRevisions {
	return v.value
}

func (v *NullableEnvironmentLatestRevisions) Set(val *EnvironmentLatestRevisions) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentLatestRevisions) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentLatestRevisions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentLatestRevisions(val *EnvironmentLatestRevisions) *NullableEnvironmentLatestRevisions {
	return &NullableEnvironmentLatestRevisions{value: val, isSet: true}
}

func (v NullableEnvironmentLatestRevisions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentLatestRevisions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
