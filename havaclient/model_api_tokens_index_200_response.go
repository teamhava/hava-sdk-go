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

// ApiTokensIndex200Response struct for ApiTokensIndex200Response
type ApiTokensIndex200Response struct {
	// The token to pass through to get the next page of results, if there are any.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// The total amount of results to be paginated
	TotalSize *int32 `json:"total_size,omitempty"`
	// The results for the current page
	Results []TokenSummary `json:"results,omitempty"`
}

// NewApiTokensIndex200Response instantiates a new ApiTokensIndex200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTokensIndex200Response() *ApiTokensIndex200Response {
	this := ApiTokensIndex200Response{}
	return &this
}

// NewApiTokensIndex200ResponseWithDefaults instantiates a new ApiTokensIndex200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTokensIndex200ResponseWithDefaults() *ApiTokensIndex200Response {
	this := ApiTokensIndex200Response{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ApiTokensIndex200Response) GetNextPageToken() string {
	if o == nil || isNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTokensIndex200Response) GetNextPageTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ApiTokensIndex200Response) HasNextPageToken() bool {
	if o != nil && !isNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ApiTokensIndex200Response) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *ApiTokensIndex200Response) GetTotalSize() int32 {
	if o == nil || isNil(o.TotalSize) {
		var ret int32
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTokensIndex200Response) GetTotalSizeOk() (*int32, bool) {
	if o == nil || isNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *ApiTokensIndex200Response) HasTotalSize() bool {
	if o != nil && !isNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given int32 and assigns it to the TotalSize field.
func (o *ApiTokensIndex200Response) SetTotalSize(v int32) {
	o.TotalSize = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ApiTokensIndex200Response) GetResults() []TokenSummary {
	if o == nil || isNil(o.Results) {
		var ret []TokenSummary
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTokensIndex200Response) GetResultsOk() ([]TokenSummary, bool) {
	if o == nil || isNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ApiTokensIndex200Response) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []TokenSummary and assigns it to the Results field.
func (o *ApiTokensIndex200Response) SetResults(v []TokenSummary) {
	o.Results = v
}

func (o ApiTokensIndex200Response) MarshalJSON() ([]byte, error) {
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

type NullableApiTokensIndex200Response struct {
	value *ApiTokensIndex200Response
	isSet bool
}

func (v NullableApiTokensIndex200Response) Get() *ApiTokensIndex200Response {
	return v.value
}

func (v *NullableApiTokensIndex200Response) Set(val *ApiTokensIndex200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTokensIndex200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTokensIndex200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTokensIndex200Response(val *ApiTokensIndex200Response) *NullableApiTokensIndex200Response {
	return &NullableApiTokensIndex200Response{value: val, isSet: true}
}

func (v NullableApiTokensIndex200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTokensIndex200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


