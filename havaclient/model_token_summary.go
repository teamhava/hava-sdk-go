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

// TokenSummary struct for TokenSummary
type TokenSummary struct {
	// The unique ID of the API token
	Id *string `json:"id,omitempty"`
	// The unique ID of the account this API token is in
	AccountId *string `json:"account_id,omitempty"`
	// The name of the API token
	Name *string `json:"name,omitempty"`
	// The last time this API token was used. Will be empty if it has never been used.
	LastUsedAt *string `json:"last_used_at,omitempty"`
}

// NewTokenSummary instantiates a new TokenSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenSummary() *TokenSummary {
	this := TokenSummary{}
	return &this
}

// NewTokenSummaryWithDefaults instantiates a new TokenSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenSummaryWithDefaults() *TokenSummary {
	this := TokenSummary{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TokenSummary) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenSummary) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TokenSummary) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TokenSummary) SetId(v string) {
	o.Id = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *TokenSummary) GetAccountId() string {
	if o == nil || isNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenSummary) GetAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *TokenSummary) HasAccountId() bool {
	if o != nil && !isNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *TokenSummary) SetAccountId(v string) {
	o.AccountId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TokenSummary) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenSummary) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TokenSummary) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TokenSummary) SetName(v string) {
	o.Name = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise.
func (o *TokenSummary) GetLastUsedAt() string {
	if o == nil || isNil(o.LastUsedAt) {
		var ret string
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenSummary) GetLastUsedAtOk() (*string, bool) {
	if o == nil || isNil(o.LastUsedAt) {
		return nil, false
	}
	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *TokenSummary) HasLastUsedAt() bool {
	if o != nil && !isNil(o.LastUsedAt) {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given string and assigns it to the LastUsedAt field.
func (o *TokenSummary) SetLastUsedAt(v string) {
	o.LastUsedAt = &v
}

func (o TokenSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.LastUsedAt) {
		toSerialize["last_used_at"] = o.LastUsedAt
	}
	return json.Marshal(toSerialize)
}

type NullableTokenSummary struct {
	value *TokenSummary
	isSet bool
}

func (v NullableTokenSummary) Get() *TokenSummary {
	return v.value
}

func (v *NullableTokenSummary) Set(val *TokenSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenSummary(val *TokenSummary) *NullableTokenSummary {
	return &NullableTokenSummary{value: val, isSet: true}
}

func (v NullableTokenSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


