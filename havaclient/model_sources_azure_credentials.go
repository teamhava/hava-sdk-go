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

// SourcesAzureCredentials The parameters required to create an Azure Credentials Source
type SourcesAzureCredentials struct {
	// The name for this source
	Name *string `json:"name,omitempty"`
	// Must be set to Azure::Credentials
	Type *string `json:"type,omitempty"`
	// The ID of the Azure subscription to import from
	SubscriptionId *string `json:"subscription_id,omitempty"`
	// The GUID representing the Active Directory Tenant
	TenantId *string `json:"tenant_id,omitempty"`
	// The Client ID for your Service Principle
	ClientId *string `json:"client_id,omitempty"`
	// The Client Secret for your Service Principle
	SecretKey *string `json:"secret_key,omitempty"`
}

// NewSourcesAzureCredentials instantiates a new SourcesAzureCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourcesAzureCredentials() *SourcesAzureCredentials {
	this := SourcesAzureCredentials{}
	return &this
}

// NewSourcesAzureCredentialsWithDefaults instantiates a new SourcesAzureCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourcesAzureCredentialsWithDefaults() *SourcesAzureCredentials {
	this := SourcesAzureCredentials{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SourcesAzureCredentials) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesAzureCredentials) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SourcesAzureCredentials) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SourcesAzureCredentials) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SourcesAzureCredentials) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesAzureCredentials) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SourcesAzureCredentials) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SourcesAzureCredentials) SetType(v string) {
	o.Type = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *SourcesAzureCredentials) GetSubscriptionId() string {
	if o == nil || isNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesAzureCredentials) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || isNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *SourcesAzureCredentials) HasSubscriptionId() bool {
	if o != nil && !isNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *SourcesAzureCredentials) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *SourcesAzureCredentials) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesAzureCredentials) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *SourcesAzureCredentials) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *SourcesAzureCredentials) SetTenantId(v string) {
	o.TenantId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SourcesAzureCredentials) GetClientId() string {
	if o == nil || isNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesAzureCredentials) GetClientIdOk() (*string, bool) {
	if o == nil || isNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SourcesAzureCredentials) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SourcesAzureCredentials) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *SourcesAzureCredentials) GetSecretKey() string {
	if o == nil || isNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesAzureCredentials) GetSecretKeyOk() (*string, bool) {
	if o == nil || isNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *SourcesAzureCredentials) HasSecretKey() bool {
	if o != nil && !isNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *SourcesAzureCredentials) SetSecretKey(v string) {
	o.SecretKey = &v
}

func (o SourcesAzureCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.SubscriptionId) {
		toSerialize["subscription_id"] = o.SubscriptionId
	}
	if !isNil(o.TenantId) {
		toSerialize["tenant_id"] = o.TenantId
	}
	if !isNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !isNil(o.SecretKey) {
		toSerialize["secret_key"] = o.SecretKey
	}
	return json.Marshal(toSerialize)
}

type NullableSourcesAzureCredentials struct {
	value *SourcesAzureCredentials
	isSet bool
}

func (v NullableSourcesAzureCredentials) Get() *SourcesAzureCredentials {
	return v.value
}

func (v *NullableSourcesAzureCredentials) Set(val *SourcesAzureCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableSourcesAzureCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableSourcesAzureCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourcesAzureCredentials(val *SourcesAzureCredentials) *NullableSourcesAzureCredentials {
	return &NullableSourcesAzureCredentials{value: val, isSet: true}
}

func (v NullableSourcesAzureCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourcesAzureCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


