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

// SourcesIndex200Response struct for SourcesIndex200Response
type SourcesIndex200Response struct {
	// The token to pass through to get the next page of results, if there are any.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// The total amount of results to be paginated
	TotalSize *int32 `json:"total_size,omitempty"`
	// The results for the current page
	Results []Source `json:"results,omitempty"`
}

// NewSourcesIndex200Response instantiates a new SourcesIndex200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourcesIndex200Response() *SourcesIndex200Response {
	this := SourcesIndex200Response{}
	return &this
}

// NewSourcesIndex200ResponseWithDefaults instantiates a new SourcesIndex200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourcesIndex200ResponseWithDefaults() *SourcesIndex200Response {
	this := SourcesIndex200Response{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *SourcesIndex200Response) GetNextPageToken() string {
	if o == nil || isNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesIndex200Response) GetNextPageTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *SourcesIndex200Response) HasNextPageToken() bool {
	if o != nil && !isNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *SourcesIndex200Response) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *SourcesIndex200Response) GetTotalSize() int32 {
	if o == nil || isNil(o.TotalSize) {
		var ret int32
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesIndex200Response) GetTotalSizeOk() (*int32, bool) {
	if o == nil || isNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *SourcesIndex200Response) HasTotalSize() bool {
	if o != nil && !isNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given int32 and assigns it to the TotalSize field.
func (o *SourcesIndex200Response) SetTotalSize(v int32) {
	o.TotalSize = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *SourcesIndex200Response) GetResults() []Source {
	if o == nil || isNil(o.Results) {
		var ret []Source
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesIndex200Response) GetResultsOk() ([]Source, bool) {
	if o == nil || isNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *SourcesIndex200Response) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Source and assigns it to the Results field.
func (o *SourcesIndex200Response) SetResults(v []Source) {
	o.Results = v
}

func (o SourcesIndex200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	if !isNil(o.TotalSize) {
		toSerialize["total_size"] = o.TotalSize
	}
	if !isNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableSourcesIndex200Response struct {
	value *SourcesIndex200Response
	isSet bool
}

func (v NullableSourcesIndex200Response) Get() *SourcesIndex200Response {
	return v.value
}

func (v *NullableSourcesIndex200Response) Set(val *SourcesIndex200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSourcesIndex200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSourcesIndex200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourcesIndex200Response(val *SourcesIndex200Response) *NullableSourcesIndex200Response {
	return &NullableSourcesIndex200Response{value: val, isSet: true}
}

func (v NullableSourcesIndex200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourcesIndex200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
