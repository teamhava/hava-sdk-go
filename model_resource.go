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

// Resource struct for Resource
type Resource struct {
	// The unique ID of the resource
	Id *string `json:"id,omitempty"`
	// The name of the resource
	Name *string `json:"name,omitempty"`
	// The type of the resource
	Type *string `json:"type,omitempty"`
	// The ID assigned to this resource by the provider
	ProviderId *string `json:"provider_id,omitempty"`
	// The estimated price per month assigned by Hava
	Price *int32 `json:"price,omitempty"`
	// The region or location in which this resource is located
	Region *string `json:"region,omitempty"`
	// Resource specific data
	Data map[string]interface{} `json:"data,omitempty"`
	// An array tags associated to this resource
	Tags []ResourceTagsInner `json:"tags,omitempty"`
	// A list of connections to other resources
	Connections []ResourceConnectionsInner `json:"connections,omitempty"`
}

// NewResource instantiates a new Resource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResource() *Resource {
	this := Resource{}
	return &this
}

// NewResourceWithDefaults instantiates a new Resource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceWithDefaults() *Resource {
	this := Resource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Resource) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Resource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Resource) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Resource) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Resource) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Resource) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Resource) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Resource) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Resource) SetType(v string) {
	o.Type = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *Resource) GetProviderId() string {
	if o == nil || isNil(o.ProviderId) {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetProviderIdOk() (*string, bool) {
	if o == nil || isNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *Resource) HasProviderId() bool {
	if o != nil && !isNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *Resource) SetProviderId(v string) {
	o.ProviderId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *Resource) GetPrice() int32 {
	if o == nil || isNil(o.Price) {
		var ret int32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetPriceOk() (*int32, bool) {
	if o == nil || isNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *Resource) HasPrice() bool {
	if o != nil && !isNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given int32 and assigns it to the Price field.
func (o *Resource) SetPrice(v int32) {
	o.Price = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Resource) GetRegion() string {
	if o == nil || isNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetRegionOk() (*string, bool) {
	if o == nil || isNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Resource) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *Resource) SetRegion(v string) {
	o.Region = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Resource) GetData() map[string]interface{} {
	if o == nil || isNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Resource) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *Resource) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Resource) GetTags() []ResourceTagsInner {
	if o == nil || isNil(o.Tags) {
		var ret []ResourceTagsInner
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetTagsOk() ([]ResourceTagsInner, bool) {
	if o == nil || isNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Resource) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTagsInner and assigns it to the Tags field.
func (o *Resource) SetTags(v []ResourceTagsInner) {
	o.Tags = v
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *Resource) GetConnections() []ResourceConnectionsInner {
	if o == nil || isNil(o.Connections) {
		var ret []ResourceConnectionsInner
		return ret
	}
	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetConnectionsOk() ([]ResourceConnectionsInner, bool) {
	if o == nil || isNil(o.Connections) {
		return nil, false
	}
	return o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *Resource) HasConnections() bool {
	if o != nil && !isNil(o.Connections) {
		return true
	}

	return false
}

// SetConnections gets a reference to the given []ResourceConnectionsInner and assigns it to the Connections field.
func (o *Resource) SetConnections(v []ResourceConnectionsInner) {
	o.Connections = v
}

func (o Resource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.ProviderId) {
		toSerialize["provider_id"] = o.ProviderId
	}
	if !isNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !isNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Connections) {
		toSerialize["connections"] = o.Connections
	}
	return json.Marshal(toSerialize)
}

type NullableResource struct {
	value *Resource
	isSet bool
}

func (v NullableResource) Get() *Resource {
	return v.value
}

func (v *NullableResource) Set(val *Resource) {
	v.value = val
	v.isSet = true
}

func (v NullableResource) IsSet() bool {
	return v.isSet
}

func (v *NullableResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResource(val *Resource) *NullableResource {
	return &NullableResource{value: val, isSet: true}
}

func (v NullableResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
